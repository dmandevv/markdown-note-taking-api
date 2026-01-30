package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("POST /upload", UploadFileHandler)
	mux.HandleFunc("GET /edit/{fileName}", EditFileHandler)
	mux.HandleFunc("POST /update", UpdateFileHandler)
	mux.HandleFunc("GET /view/{fileName}", ViewFileHandler)
	http.ListenAndServe(":8080", mux)
}
