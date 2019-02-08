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

	response, err := json.Marshal(charCount(data.Data))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func charCount(s string) map[string]int {
	var myMap map[string]int
	myMap = make(map[string]int)
	charnum := len(s)
	char := strings.Split(s, "")

	myMap[char[0]] = charnum
	return myMap
}
