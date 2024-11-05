package utils

import (
	"fmt"
	"log"

	"github.com/stellar/go/keypair"
)

func KeyPair() (string, string){
	pair, err := keypair.Random()

	if err != nil {
		log.Fatalf(fmt.Sprintf("An error occurred while generating key pair: %v ", err.Error()))
	}

	privateKey := pair.Seed()
	publicKey := pair.Address()
	

	log.Println("Private Key: ")
	log.Println(privateKey)

	log.Printf("Public Key: ")
	log.Println(publicKey)

	return privateKey, publicKey
}