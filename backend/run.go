package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

//Cmd will define a bash command
type Cmd struct {
	Command string `json:"command"`
}

func handleRun(w http.ResponseWriter, r *http.Request) {
	command := Cmd{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		log.Fatal(err)
	}

	response, err := json.Marshal(runCmd(command.Command))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func runCmd(c string) string {
	out, err := exec.Command(c).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
