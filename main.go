package main

import (
	"html/template"
	"net/http"

	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
)

/* multiplexer - piece of code that redirects a request to a
 * handler. net/http libs provide a default multiplexer that
 * can be created by calling the NewServeMux function
 */

func main() {

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// handle redirect of root URL to a handler finction
	mux.HandleFunc("/", index)
	// handle error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// start the server
	server := &http.Server{
		Addr:    "0.0.0.0.0:8000",
		Handler: mux,
	}
	server.ListenAndServe()

}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
