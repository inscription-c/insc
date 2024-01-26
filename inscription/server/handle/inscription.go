package handle

import (
	"errors"
	"github.com/btcsuite/btcd/wire"
	"github.com/gin-gonic/gin"
	"github.com/inscription-c/insc/constants"
	"github.com/inscription-c/insc/inscription/index"
	"github.com/inscription-c/insc/inscription/index/tables"
	"github.com/inscription-c/insc/internal/util"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

// RespInscription is a struct that represents the response for an inscription request.
type RespInscription struct {
	InscriptionId   string   `json:"inscription_id"`
	Charms          []string `json:"charms"`
	InscriptionNum  int64    `json:"inscription_number"`
	Next            string   `json:"next"`
	Previous        string   `json:"previous"`
	Address         string   `json:"address"`
	Sat             uint64   `json:"sat"`
	ContentLength   int      `json:"content_length"`
	ContentType     string   `json:"content_type"`
	GenesisFee      uint64   `json:"genesis_fee"`
	GenesisHeight   uint32   `json:"genesis_height"`
	OutputValue     int64    `json:"output_value"`
	SatPoint        string   `json:"satpoint"`
	Timestamp       int64    `json:"timestamp"`
	DstChain        string   `json:"dst_chain"`
	ContentProtocol string   `json:"content_protocol"`
}

// Inscription is a handler function for handling inscription requests.
// It validates the request parameters and calls the doInscription function.
func (h *Handler) Inscription(ctx *gin.Context) {
	query := ctx.Param("query")
	if query == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if err := h.doInscription(ctx, query); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
}

// doInscription is a helper function for handling inscription requests.
// It retrieves the inscription based on the provided query and returns it in the response.
func (h *Handler) doInscription(ctx *gin.Context, query string) error {
	// Trim spaces from the query and try to convert it to an inscription ID.
	// If that fails, try to convert it to a sequence number.
	// If both fail, return a bad request status.
	query = strings.TrimSpace(query)
	inscriptionId := util.StringToInscriptionId(query)
	var err error
	var inscription tables.Inscriptions
	if inscriptionId != nil {
		inscription, err = h.DB().GetInscriptionById(inscriptionId.String())
		if err != nil {
			return err
		}
	} else {
		var sequenceNum uint64
		sequenceNum, err = strconv.ParseUint(query, 10, 64)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return nil
		}
		inscription, err = h.DB().GetInscriptionBySequenceNum(sequenceNum)
		if err != nil {
			return err
		}
	}

	// If the inscription does not exist, return a not found status.
	if inscription.Id == 0 {
		ctx.Status(http.StatusNotFound)
		return nil
	}

	// Retrieve the previous and next inscriptions.
	preInscription, err := h.DB().GetInscriptionBySequenceNum(inscription.SequenceNum - 1)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	preInscriptionId := ""
	if preInscription.Id > 0 {
		preInscriptionId = preInscription.Outpoint.InscriptionId().String()
	}

	nextInscription, err := h.DB().GetInscriptionBySequenceNum(inscription.SequenceNum + 1)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	nextInscriptionId := ""
	if nextInscription.Id > 0 {
		nextInscriptionId = nextInscription.Outpoint.InscriptionId().String()
	}

	// Retrieve the value of the inscription's outpoint.
	value, err := h.DB().GetValueByOutpoint(inscription.Outpoint.String())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Retrieve the transaction of the inscription's outpoint and get the taproot address.
	tx, err := h.RpcClient().GetRawTransaction(&inscription.Outpoint.Hash)
	if err != nil {
		return err
	}
	pkScript := tx.MsgTx().TxOut[inscription.Outpoint.Index].PkScript
	taprootAddress, err := util.TapRootAddress(pkScript, h.GetChainParams())
	if err != nil {
		return err
	}

	// Retrieve the sat point of the inscription.
	satPoint, err := h.DB().GetSatPointBySat(inscription.Sat)
	if err != nil {
		return err
	}
	satPointStr := util.FormatSatPoint(wire.OutPoint{}.String(), 0)
	if satPoint.Id == 0 {
		satPointStr = util.FormatSatPoint(satPoint.Outpoint, satPoint.Offset)
	}

	// Check if the inscription's body is a BRC20C token.
	brc20c := &util.BRC20C{}
	brc20c.Reset(inscription.Body)
	contentProtocol := ""
	if brc20c.Check() == nil {
		contentProtocol = constants.ProtocolBRC20C
	}

	// Create the response and return it.
	resp := &RespInscription{
		InscriptionId:   inscriptionId.String(),
		InscriptionNum:  inscription.InscriptionNum,
		Charms:          index.CharmsAll.Titles(inscription.Charms),
		GenesisHeight:   inscription.Height,
		GenesisFee:      inscription.Fee,
		OutputValue:     value,
		Address:         taprootAddress,
		Sat:             inscription.Sat,
		SatPoint:        satPointStr,
		ContentType:     inscription.ContentType,
		ContentLength:   len(inscription.Body),
		Timestamp:       inscription.Timestamp,
		DstChain:        inscription.DstChain,
		ContentProtocol: contentProtocol,
		Previous:        preInscriptionId,
		Next:            nextInscriptionId,
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}
