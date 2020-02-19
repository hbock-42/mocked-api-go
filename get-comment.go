package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Comments contains an array of Comment
type Comments struct {
	Comments []Comment `json:"comments"`
}

// Comment related to a Post
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// GetComment returns all Comments linked to a postID
func GetComment(postID int) ([]Comment, error) {
	jsonFile, err := os.Open("./data/posts.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var allComments Comments
	json.Unmarshal(byteValue, &allComments)

	comments := []Comment{}
	for _, comment := range allComments.Comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments, nil
}
