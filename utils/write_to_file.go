package utils

import (
	"fmt"
	"io/ioutil"
)

func WriteToFile(content string) error{
	if(len(content) < 4){
		return fmt.Errorf("content must be at least 4 characters long")
	}

	fileName := content[:4] + ".txt"

	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil{
		return err
	}

	fmt.Printf("Content written to file: %s\n", fileName)
	return nil
}