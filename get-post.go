package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Posts contains an array of Post
type Posts struct {
	Posts []Post `json:"posts"`
}

// Post from a user
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// GetPost returns a comment from id
func GetPost(id int) (*Post, error) {
	jsonFile, err := os.Open("./data/posts.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var posts Posts
	json.Unmarshal(byteValue, &posts)
	// dirty, we assume that the data is ordered
	if len(posts.Posts) > id-1 {
		return &(posts.Posts[id-1]), nil
	}
	return nil, errors.New("Unable to find a post with id " + string(id))
}
