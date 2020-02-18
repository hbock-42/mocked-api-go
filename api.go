package main

//import two built-in Go packages to use for building our application
import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	//assign a handler to the root path
	http.HandleFunc("/", roothandler)
	http.HandleFunc("/posts/", saluthandler)

	//listen on port 8080 for incoming http requests
	http.ListenAndServe(":8080", nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	//write a message to the http response
	fmt.Fprintf(w, "Welcome to the simplest http service")
}

func saluthandler(w http.ResponseWriter, r *http.Request) {
	urlChunks := strings.Split(r.URL.Path, "/")
	if len(urlChunks) > 2 && urlChunks[2] != "" {
		postID, err := strconv.Atoi(urlChunks[2])
		if err != nil {
			fmt.Println(err)
		} else {
			GetPost(postID)
		}
	}
}
