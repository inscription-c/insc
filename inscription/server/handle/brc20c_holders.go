package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/inscription-c/insc/constants"
	"github.com/inscription-c/insc/internal/util"
	"net/http"
)

func (h *Handler) BRC20CHolders(ctx *gin.Context) {
	tkid := ctx.Query("tkid")
	page := ctx.DefaultQuery("page", "1")
	if tkid == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if gconv.Int(page) < 1 {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if err := h.doBRC20CHolders(ctx, tkid, gconv.Int(page)); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) doBRC20CHolders(ctx *gin.Context, tkid string, page int) error {
	pageSize := 100
	inscriptionId := util.StringToInscriptionId(tkid)
	if inscriptionId == nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	protocol, err := h.DB().GetProtocolByOutpoint(inscriptionId.OutPoint.String())
	if err != nil {
		return err
	}
	if protocol.Id == 0 || protocol.Protocol != constants.ProtocolBRC20C {
		ctx.Status(http.StatusNotFound)
		return nil
	}
	if protocol.Operator == constants.OperationMint {
		protocol, err = h.DB().GetProtocolByOutpoint(protocol.TkID)
		if err != nil {
			return err
		}
		if protocol.Id == 0 {
			ctx.Status(http.StatusNotFound)
			return nil
		}
	}

	list, err := h.DB().FindHoldersByTkId(protocol.Outpoint.String(), constants.ProtocolBRC20C, constants.OperationMint, page, pageSize)
	if err != nil {
		return err
	}
	more := false
	if len(list) > pageSize {
		more = true
		list = list[:pageSize]
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page_index": page,
		"more":       more,
		"holders":    list,
	})
	return nil
}
