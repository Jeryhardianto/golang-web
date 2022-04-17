package main

import (
	"golang-web/handler"
	"log"
	"net/http"
)

func main() {
	// Mux
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/hello", handler.Sayhello) 
	mux.HandleFunc("/about", handler.About) 
	mux.HandleFunc("/post", handler.PostForm) 
	mux.HandleFunc("/process", handler.Process) 

	// Post and Get
	mux.HandleFunc("/post-get", handler.PostGet)

	//Load file static / public
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Yang terjadi dibelakang layar
	// localhost:8080/static/style.css



	// Routes Query String
	mux.HandleFunc("/product", handler.Product)

	// Start server
	log.Printf("Starting server on port 9999")
	err := http.ListenAndServe(":9999", mux)
	log.Fatal(err)
	
}

