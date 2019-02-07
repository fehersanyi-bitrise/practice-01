package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	//tbd
	log.Print(r.Body)
	log.Print(w.Header())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", handleFizzBuzz).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", nil))

}

func counter(n int) string {
	if n%15 == 0 {
		return "fizzbuzz"
	} else if n%5 == 0 {
		return "buzz"
	} else if n%3 == 0 {
		return "fizz"
	} else {
		return string(n)
	}
}
