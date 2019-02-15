package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleAPIGetLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := command{}
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		log.Printf("Decoder: %s", err)
	}
	mds, err := callDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	resp, err := json.Marshal(mds[cmd.ID].Log)
	if err != nil {
		log.Println(err)
	}
	w.Write(resp)
}
