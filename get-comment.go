package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// CommentsHandler regroup methods thats handle comments request
type CommentsHandler struct {
}

// Handler who takes a post id as arguments
func (h *CommentsHandler) Handler(postID int) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			comments, err := h.GetComments(postID)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			json, err := json.Marshal(comments)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			res.Write(json)
		default:
			http.Error(res, "Only GET is allowed", http.StatusMethodNotAllowed)
		}
	})
}

// GetComments returns all Comments linked to a postID
func (h *CommentsHandler) GetComments(postID int) ([]Comment, error) {
	jsonFile, err := os.Open("./data/comments.json")
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
