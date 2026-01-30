package main

import (
	"html/template"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./assets/uploads")
	if err != nil {
		http.Error(w, "Failed to read files directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	tmpl, err := template.ParseFiles("./assets/index.html")
	if err != nil {
		http.Error(w, "Failed to parse index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, fileNames)
	if err != nil {
		http.Error(w, "Failed to insert file names into INDEX page: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
