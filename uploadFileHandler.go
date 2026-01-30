package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	r.ParseMultipartForm(10485760)                         //10MB max file size
	uploadedFile, handler, err := r.FormFile("fileUpload") //form field must be "fileUpload"
	if err != nil {
		http.Error(w, "Failed to upload file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer uploadedFile.Close()

	//create empty file on server
	serverFilePath := filepath.Join("./assets/uploads", handler.Filename)
	serverFile, err := os.Create(serverFilePath)
	if err != nil {
		http.Error(w, "Failed to create file on server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer serverFile.Close()

	//copy uploaded files data to empty server file
	_, err = io.Copy(serverFile, uploadedFile)
	if err != nil {
		http.Error(w, "Failed to save uploaded files data to server file"+err.Error(), http.StatusInternalServerError)
		return
	}

	//w.Write([]byte(serverFilePath))
	http.Redirect(w, r, fmt.Sprintf("/view/%s", handler.Filename), http.StatusSeeOther)
}
