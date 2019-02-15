package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/API/run", handleRun).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}

type command struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
	Flag    string `json:"flag"`
	Log     string `json:"log"`
}

func handleRun(w http.ResponseWriter, r *http.Request) {
	command := command{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		log.Fatal(err)
	}
	exec, err := runCmd(command.Command, command.Flag)
	if err != nil {
		log.Printf("failed to execute command %s", err)
	}
	response, err := json.Marshal(exec)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sent my response: %s", response)
	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func runCmd(c string, f ...string) (string, error) {
	if len(f) > 0 {
		if f[0] == "" {
			out, err := exec.Command(c).Output()
			if err != nil {
				return "", err
			}
			return string(out), nil
		}
	}
	out, err := exec.Command(c, f[0]).Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
