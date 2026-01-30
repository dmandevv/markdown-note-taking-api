package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func EditFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	filePath := filepath.Join("./assets/uploads", r.PathValue("fileName"))
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to read file: %s, %w", filePath, err).Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(fileData)

}
