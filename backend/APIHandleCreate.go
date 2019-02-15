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
		log.Panicln(err)
	}
	defer db.Close()

	_, err = db.Query(`insert into cmdb (ID, COMMAND, LOGS) values ($1, $2, $3)`, cmd.ID, cmd.Command, cmd.Log)
	if err != nil {
		log.Panicln(err)
	}

	response, err := json.Marshal(cmd)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("%#v", response)
	if _, err := w.Write(response); err != nil {
		log.Panicln(err)
	}
}
