package main

//import two built-in Go packages to use for building our application
import (
	"fmt"
	"net/http"
)

func main() {
	//assign a handler to the root path
	http.HandleFunc("/", roothandler)
	http.HandleFunc("/salut", saluthandler)

	//listen on port 8080 for incoming http requests
	http.ListenAndServe(":8080", nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	//write a message to the http response
	fmt.Fprintf(w, "Welcome to the simple http service")
}

func saluthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola")
}
