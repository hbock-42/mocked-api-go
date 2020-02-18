package main

import (
	"fmt"
	"os"
	// "encoding/json"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// GetPost returns a comment from id
func GetPost(id int) {
	jsonFile, err := os.Open("./data/posts.json")
	if err != nil {
		fmt.Println(err)
	} else {

	}
	defer jsonFile.Close()
}
