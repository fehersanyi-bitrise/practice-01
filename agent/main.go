package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/API/run", runCommand).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func runCommand(w http.ResponseWriter, r *http.Request) {

}
