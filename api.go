package main

import (
	"net/http"
	"path"
	"strings"
)

// tried this for routing
// https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html

func main() {
	a := &App{
		PostsHandler: new(PostsHandler),
	}
	http.ListenAndServe(":8000", a)
}

// App is entry handler
type App struct {
	PostsHandler *PostsHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "posts" {
		h.PostsHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
