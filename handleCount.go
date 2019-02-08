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

	response, err := json.Marshal(dupCount(data.Data))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func dupCount(s string) map[string]int {
	list := strings.Split(s, "")
	duplicateFrequency := make(map[string]int)

	for _, item := range list {

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item]++
		} else {
			duplicateFrequency[item] = 1
		}
	}
	return duplicateFrequency
}
