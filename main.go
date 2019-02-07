package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type num struct {
	Number int `json:"number"`
}

//Response is a res
type Response struct {
	Response string `json:"response"`
}

func handleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	n := num{}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		log.Fatal(err)
	}
	//log.Print(n.Number)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Response{}
	resp.Response = counter(n.Number)
	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", handleFizzBuzz).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", router))

}

func counter(n int) string {
	if n%15 == 0 {
		return "fizzbuzz"
	} else if n%5 == 0 {
		return "buzz"
	} else if n%3 == 0 {
		return "fizz"
	} else {
		return strconv.Itoa(n)
	}
}
