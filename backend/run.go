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
	exec, err := runCmd(command.Command)
	if err != nil {
		log.Printf("failed to execute command %s", err)
	}
	response, err := json.Marshal(exec)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func runCmd(c string) (string, error) {
	out, err := exec.Command(c).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
