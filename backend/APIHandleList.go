package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func handleAPIGetCommands(w http.ResponseWriter, r *http.Request) {

	cmd, err := callDB()
	if err != nil {
		log.Println(err)
	}
	res, err := json.Marshal(cmd)
	if err != nil {
		log.Println(err)
	}
	w.Write(res)
}

func callDB() ([]command, error) {
	cmd := []command{}
	connstr := "user=infra password=1234 dbname=infra host=db sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return []command{}, err
	}
	defer db.Close()

	rows, err := db.Query("select * from cmdb")
	if err != nil {
		return []command{}, err
	}
	for rows.Next() {
		c := command{}
		if err = rows.Scan(&c.ID, &c.Command, &c.Log); err != nil {
			return []command{}, err
		}
		cmd = append(cmd, c)
	}
	return cmd, nil
}
