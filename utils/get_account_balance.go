package utils

import (
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon"
)

func AccountBalance(publicKey string)  []horizon.Balance {
	request := horizonclient.AccountRequest{AccountID: publicKey}

	//TODO: change to public network
	account, err := horizonclient.DefaultTestNetClient.AccountDetail(request)

	if err != nil{
		log.Fatal(err)
	}

	log.Println("Balances for account: ", publicKey)

	for _, balance := range account.Balances{
		log.Println(balance)
	}

	return account.Balances
}