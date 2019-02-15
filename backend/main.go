package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", HandleFizzBuzz).Methods("GET")
	router.HandleFunc("/count", handleCount).Methods("POST")
	router.HandleFunc("/run", handleRun).Methods("POST")
	router.HandleFunc("/API/list", handleAPIGetCommands).Methods("GET")
	router.HandleFunc("/API/get", handleAPIGetLog).Methods("GET")
	router.HandleFunc("/API/create", APIHandleCreate).Methods("POST")
	log.Fatal(http.ListenAndServe(":3030", router))
}
