package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FundAccount(publicKey string) string{
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + publicKey)

	if err != nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(body))

	return string(body)
}