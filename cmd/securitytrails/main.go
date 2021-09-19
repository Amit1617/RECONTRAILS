package main

import (
	"fmt"
	"os"
	"log"

	"github.com/encodedguy/recontrails/securitytrails"
)

func main() {
	if len(os.Args) != 2{
		log.Fatalln("Usage: ./main")
	}

	apiKey := os.Getenv("SECURITYTRAILS_API_KEY")

	s := securitytrails.New(apiKey)
	info, err := s.APIInfo()

	if err != nil{
		log.Panicln(err)
	}

	fmt.Printf(
		"APIKey Valid: %t", info.Success
	)
}
