package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

type ViewPageData struct {
	FileName      string
	FormattedText string
}

func ViewFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fileName := r.PathValue("fileName")
	filepath := filepath.Join("./assets/uploads", fileName)
	rawFileData, err := os.ReadFile(filepath)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to find file: %s, %w", fileName, err).Error(), http.StatusNotFound)
		return
	}

	//Convert raw markdown bytes into html format
	//var buf bytes.Buffer
	if err = goldmark.Convert(rawFileData, w); err != nil {
		http.Error(w, "Failed to convert raw markdown file to HTML: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/*

		//prepare template to upload data to form
		tmpl, err := template.ParseFiles("./assets/view.html")
		if err != nil {
			http.Error(w, "Failed to parse VIEW page: "+err.Error(), http.StatusInternalServerError)
			return
		}

		viewPageData := ViewPageData{
			FileName:      fileName,
			FormattedText: buf.String(),
		}

		err = tmpl.Execute(w, viewPageData)
		if err != nil {
			http.Error(w, "Failed to insert text data into VIEW page :"+err.Error(), http.StatusInternalServerError)
			return
		}
	*/

}
