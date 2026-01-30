package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func UpdateFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	err := r.ParseMultipartForm(10485760) //10MB max file size
	if err != nil {
		http.Error(w, "Failed to parse file update form: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fileName := r.FormValue("fileName")
	rawText := r.FormValue("rawText")

	filepath := filepath.Join("./assets/uploads", fileName)

	fileInfo, err := os.Stat(filepath)
	if err != nil {
		http.Error(w, "Failed to find file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(filepath, []byte(rawText), fileInfo.Mode().Perm())
	if err != nil {
		http.Error(w, "Failed to update file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/view/%s", fileName), http.StatusSeeOther)

}
