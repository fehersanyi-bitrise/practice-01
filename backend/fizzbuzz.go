package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type num struct {
	Number int `json:"number"`
}

type res struct {
	Response string `json:"response"`
}

//HandleFizzBuzz ...
func HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	num := num{}
	resp := res{}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&num); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(handleBadReqest())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		resp.Response = counter(num.Number)
		response, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := w.Write(response); err != nil {
			log.Fatal(err)
		}
	}
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
