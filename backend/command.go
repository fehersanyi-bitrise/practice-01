package main

type command struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Log     string `json:"log"`
}
