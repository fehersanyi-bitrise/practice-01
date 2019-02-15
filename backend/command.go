package main

type command struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
	Flag    string `json:"flag"`
	Log     string `json:"log"`
}
