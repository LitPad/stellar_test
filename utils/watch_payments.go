package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon/operations"
)

func Watch(publicKey string){
	client := horizonclient.DefaultTestNetClient
	opRequest := horizonclient.OperationRequest{ForAccount: publicKey, Cursor: "now", IncludeFailed: true}
	
	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		time.Sleep(60 * time.Second)
		cancel()
	}()

	pribtHander := func (op operations.Operation)  {
		fmt.Print("Incoming payment: ")
		fmt.Println(op)
		// store in the db
		fmt.Print("----")
	}

	err := client.StreamPayments(ctx, opRequest, pribtHander)

	if err != nil{
		fmt.Println(err)
	}
}