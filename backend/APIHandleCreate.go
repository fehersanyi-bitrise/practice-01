package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

//APIHandleCreate inserts a new command into the database
func APIHandleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := command{}
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	connstr := "user=infra password=1234 dbname=infra host=db sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	cmd.Log, err = runCmd(cmd.Command)
	if err != nil {
		cmd.Log = err.Error()
	}
	_, err = db.Query(`insert into cmdb (COMMAND, LOGS) values ($1, $2)`, cmd.Command, cmd.Log)
	if err != nil {
		log.Printf("Error adding row: %s", err)
	}

	response, err := json.Marshal(cmd.Log)
	if err != nil {
		log.Println(err)
	}

	if _, err := w.Write(response); err != nil {
		log.Println(err)
	}
}
