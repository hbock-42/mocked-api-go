package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Posts contains an array of Post
type Posts struct {
	Posts []Post `json:"posts"`
}

// Post from a user
type Post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	MinSalary int    `json:"min_salary"`
}

// PostsHandler regroup methods thats handle posts request
type PostsHandler struct {
	CommentsHandler *CommentsHandler
}

func (h *PostsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid post id %q", head), http.StatusBadRequest)
		return
	}
	if req.URL.Path != "/" {
		head, _ := ShiftPath(req.URL.Path)
		switch head {
		case "comments":
			h.CommentsHandler.Handler(id).ServeHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
		return
	}

	switch req.Method {
	case "GET":
		posts, err := h.GetPost(id)
		if err != nil {
		} else {
			json, err := json.Marshal(posts)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			res.Write(json)
		}
	default:
		http.Error(res, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

// GetPost returns a comment from id
func (h *PostsHandler) GetPost(id int) (*Post, error) {
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
