package inscription

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/dotbitHQ/insc/config"
	"github.com/dotbitHQ/insc/index"
	"github.com/dotbitHQ/insc/internal/signal"
	"github.com/dotbitHQ/insc/wallet"
	"github.com/spf13/cobra"
	"os"
)

// InsufficientBalanceError is an error that represents an insufficient balance.
var InsufficientBalanceError = errors.New("InsufficientBalanceError")

// Cmd is a cobra command that runs the inscribe function when executed.
// It also handles any errors returned by the inscribe function.
var Cmd = &cobra.Command{
	Use:   "inscribe",
	Short: "inscription casting",
	Run: func(cmd *cobra.Command, args []string) {
		if err := inscribe(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		signal.SimulateInterrupt()
		<-signal.InterruptHandlersDone
	},
}

// inscribe is a function that performs the inscription process.
// It checks the configuration, gets the UTXO, creates the commit and reveal transactions, and signs and sends the transactions.
// It also handles any errors that occur during these processes.
func inscribe() error {
	// Check the configuration
	if err := configCheck(); err != nil {
		return err
	}

	// Get the database
	db := index.DB()
	signal.AddInterruptHandler(func() {
		if err := db.Close(); err != nil {
			log.Error("db.Close", "err", err)
		}
	})

	// Create a new wallet client
	walletCli, err := wallet.NewWalletClient(
		config.RpcConnect,
		config.Username,
		config.Password,
		&rpcclient.NotificationHandlers{
			OnClientConnected: OnClientConnected,
		},
	)
	if err != nil {
		return err
	}
	signal.AddInterruptHandler(func() {
		walletCli.Shutdown()
	})

	// Create a new inscription from the file path
	inscription, err := NewFromPath(config.FilePath,
		WithWalletClient(walletCli),
		WithPostage(config.Postage),
		WithDstChain(config.DstChain),
		WithWalletPass(config.WalletPass),
		WithCborMetadata(config.CborMetadata),
		WithJsonMetadata(config.JsonMetadata),
	)
	if err != nil {
		return err
	}

	// Get all UTXO for all unspent addresses and exclude the UTXO where the inscription
	if err := inscription.getUtxo(); err != nil {
		return err
	}

	// Create commit and reveal transactions
	if err := inscription.CreateInscriptionTx(); err != nil {
		return err
	}

	// If it's a dry run, log the success and the transaction IDs and return
	if config.DryRun {
		log.Info("dry run success")
		log.Info("commitTx: ", inscription.CommitTxId())
		log.Info("revealTx: ", inscription.RevealTxId())
		return nil
	}

	// Sign the reveal transaction
	if err := inscription.SignRevealTx(); err != nil {
		return err
	}
	// Sign the commit transaction
	if err := inscription.SignCommitTx(); err != nil {
		return err
	}

	// Send the commit transaction
	commitTxHash, err := walletCli.SendRawTransaction(inscription.commitTx, false)
	if err != nil {
		return err
	}
	log.Info("commitTxSendSuccess", commitTxHash)

	// Send the reveal transaction
	revealTxHash, err := walletCli.SendRawTransaction(inscription.revealTx, false)
	if err != nil {
		return err
	}
	log.Info("revealTxSendSuccess", revealTxHash)

	return nil
}
