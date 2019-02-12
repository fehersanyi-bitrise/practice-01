package main

import (
	"encoding/json"
	"log"
)

func handleBadReqest() []byte {
	errMessage := "Bad request"
	errResponse, err := json.Marshal(errMessage)
	if err != nil {
		log.Fatal(err)
	}
	return errResponse
}
