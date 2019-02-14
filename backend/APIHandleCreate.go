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
	//	initDB(db)

	connStr := "user=infra password=1234 dbname=infra sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to DB %s\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing connection: %s\n", err)
		}
	}()

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	c := command{}
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Fatalf("Error decoding request %s\n", err)
	}

	rows, err := db.Query("SELECT * FROM cmdb")
	if err != nil {
		log.Fatalf("Error querrying table %s\n", err)
	}

	rv := command{}
	for rows.Next() {
		if err = rows.Scan(&rv.ID, &rv.Command, &rv.Log); err != nil {
			log.Fatalf("Error converting values %s\n", err)
		}
	}

	// res, err := db.Query(`INSERT INTO cmdb (ID, COMMAND, LOGS) VALUES ($1, $2, $3);`, c.ID, c.Command, c.Log)
	// send, err := res.Columns()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if rv.ID == "" {
		rv.ID = "empty"
	}
	if _, err := w.Write([]byte(rv.ID)); err != nil {
		log.Fatalf("Error sending back results %s\n", err)
	}
}
