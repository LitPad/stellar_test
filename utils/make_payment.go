package utils

import (
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func Pay(sendersPublicKey, receiversPrivateKey, amount, purpose string){
	client := horizonclient.DefaultTestNetClient // Todo: change to public network

	destAccountRequest := horizonclient.AccountRequest{AccountID: receiversPrivateKey}

	destAccount, err := client.AccountDetail(destAccountRequest)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Destination Account: ", destAccount)

	// Load source account
	sourceKP := keypair.MustParseFull(sendersPublicKey)
	sourceAccountRequest := horizonclient.AccountRequest{AccountID: sourceKP.Address()}
	sourceAccount, err := client.AccountDetail(sourceAccountRequest)

	if err != nil{
		log.Fatal(err)
	}

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount: &sourceAccount,
			IncrementSequenceNum: true,
			BaseFee: txnbuild.MinBaseFee,
			Preconditions: txnbuild.Preconditions{
				TimeBounds: txnbuild.NewTimeout(600),
			},
			Operations: []txnbuild.Operation{
				&txnbuild.Payment{
					Destination: receiversPrivateKey,
					Amount: amount,
					Asset: txnbuild.NativeAsset{},
				},
			},
			Memo: txnbuild.MemoText(purpose),
		},
	)

	if err != nil{
		log.Fatal(err)
	}
	
	// Sign the transaction
	tx, err = tx.Sign(network.TestNetworkPassphrase, sourceKP)
	if err != nil{
		log.Fatal(err)
	}

	// Save tx to db

	resp, err := horizonclient.DefaultTestNetClient.SubmitTransaction(tx)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Successful transaction: ")
	fmt.Println("ledger: ", resp.Ledger)
	fmt.Println("hash: ", resp.Hash)
}