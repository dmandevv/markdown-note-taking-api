package main

import (
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./assets"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("POST /upload", UploadFileHandler)
	mux.HandleFunc("GET /edit/{fileName}", EditFileHandler)
	mux.HandleFunc("POST /update", UpdateFileHandler)
	mux.HandleFunc("GET /view/{fileName}", ViewFileHandler)
	http.ListenAndServe(":8080", mux)
}
