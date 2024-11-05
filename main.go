package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/litpad/stellar-wallet/utils"
)

// func main(){
// 	secretKey, address := utils.KeyPair()

// 	err := utils.WriteToFile(secretKey)

// 	if err != nil{
// 		log.Fatalf("Error writing secret key to file: %v", err)
// 	}

// 	// fund with 10,000 native tokens on testnet
// 	utils.FundAccount(address)

// 	// logs account balance
// 	utils.AccountBalance(address)
// }

func main(){
	fmt.Println("Enter address: ")

	scanner := bufio.NewScanner(os.Stdin)

	var address string

	if scanner.Scan(){
		address = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	utils.AccountBalance(address)
}