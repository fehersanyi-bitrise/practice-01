package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type data struct {
	Data string `json:"data"`
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	data := data{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	response, err := json.Marshal(dupCount(strings.Split(data.Data, "")))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func dupCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}
