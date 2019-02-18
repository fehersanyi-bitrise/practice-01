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
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

func callDB() (map[int]command, error) {
	cmdMap := make(map[int]command)
	connstr := "user=infra password=1234 dbname=infra host=db sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return cmdMap, err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	rows, err := db.Query("select * from cmdb")
	if err != nil {
		return cmdMap, err
	}
	for rows.Next() {
		c := command{}
		if err = rows.Scan(&c.ID, &c.Command, &c.Flag, &c.Log); err != nil {
			return cmdMap, err
		}
		cmdMap[c.ID] = c
	}
	return cmdMap, nil
}
