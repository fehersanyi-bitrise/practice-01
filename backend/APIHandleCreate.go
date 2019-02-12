package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type command struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Log     string `json:"log"`
}

// func initDB(db *sql.DB) {
// 	connStr := "user=infra dbname=infra sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func handleAPICreateCommand(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	//	initDB(db)

	connStr := "user=infra dbname=infra sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	c := command{}
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM cmdb;")
	if err != nil {
		log.Fatal(err)
	}

	rv := command{}
	err = res.Scan(rv.ID, rv.Command, rv.Log)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(rv.ID)
	// res, err := db.Query(`INSERT INTO cmdb (ID, COMMAND, LOGS) VALUES ($1, $2, $3);`, c.ID, c.Command, c.Log)
	// send, err := res.Columns()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	w.Write([]byte(rv.ID))
}
