package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", HandleFizzBuzz).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", router))
}
