# Markdown Note-Taking API

A modern, minimalist note-taking application that lets users upload, edit, and view markdown files with beautiful HTML rendering.

[![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Status](https://img.shields.io/badge/status-active-brightgreen)]()

## Table of Contents

- [About](#-about)
- [Features](#-features)
- [Getting Started](#-getting-started)
- [Installation](#-installation)
- [Usage](#-usage)
- [Project Structure](#-project-structure)
- [API Endpoints](#-api-endpoints)
- [Contributing](#-contributing)
- [License](#-license)

## 🚀 About

**Markdown Note-Taking API** is a lightweight web application built with Go that provides a seamless experience for managing markdown notes. This project is an exercise from the roadmap.sh project: [Markdown Note-taking App](https://roadmap.sh/projects/markdown-note-taking-app) 📚


## ✨ Features

- **📝 Markdown Upload** - Upload markdown files (.md format) with a simple drag-and-drop interface
- **✏️ Live Editing** - Edit your markdown content directly in the browser with a spacious editor
- **👁️ HTML Rendering** - Automatically render markdown to beautiful HTML with proper formatting


### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/dmandevv/markdown-note-taking-api.git
   cd markdown-note-taking-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   go build
   ```

4. **Run the application**
   ```bash
   go run .
   ```

The application will start on `http://localhost:8080` (or your configured port).

## 📖 Usage

1. **Open the Application** - Navigate to `http://localhost:8080` in your web browser
2. **Upload a File** - Click on the "Upload File" section and select a markdown (.md) file from your computer
3. **View Your Files** - All uploaded files will appear in the "Your Files" section on the home page
4. **View Content** - Click the "View" button to see your markdown rendered as HTML
5. **Edit Content** - Click the "Edit" button to modify your markdown content
6. **Save Changes** - After editing, click "Save" to persist your changes

## 📁 Project Structure

```
.
├── main.go                    # Application entry point
├── homeHandler.go             # Home page handler
├── uploadFileHandler.go       # File upload handler
├── viewFileHandler.go         # File viewing handler
├── editFileHandler.go         # File editing handler
├── updateFileHandler.go       # File update handler
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies lock file
├── assets/
│   ├── index.html             # Home page template
│   ├── view.html              # File viewer template
│   ├── edit.html              # File editor template
│   ├── style.css              # Global styles
│   └── uploads/               # Uploaded markdown files directory
└── README.md                  # This file
```

## 🔌 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Home page - lists all uploaded files |
| POST | `/upload` | Upload a new markdown file |
| GET | `/view/{filename}` | View rendered markdown file |
| GET | `/edit/{filename}` | Open file for editing |
| POST | `/update` | Save edited markdown file |


## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.