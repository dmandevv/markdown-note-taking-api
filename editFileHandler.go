package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type EditPageData struct {
	FileName string
	RawText  string
}

func EditFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fileName := r.PathValue("fileName")
	filePath := filepath.Join("./assets/uploads", fileName)
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to find file: %s, %w", fileName, err).Error(), http.StatusNotFound)
		return
	}

	//prepare template for uploading raw text to form
	tmpl, err := template.ParseFiles("./assets/edit.html")
	if err != nil {
		http.Error(w, "Failed to parse EDIT page: "+err.Error(), http.StatusInternalServerError)
		return
	}

	editPageData := EditPageData{
		FileName: fileName,
		RawText:  string(fileData),
	}

	err = tmpl.Execute(w, editPageData)
	if err != nil {
		http.Error(w, "Failed to insert text data into EDIT form: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
