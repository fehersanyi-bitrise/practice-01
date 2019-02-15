package main

type command struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
	Log     string `json:"log"`
}
