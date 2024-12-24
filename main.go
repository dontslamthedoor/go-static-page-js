package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse the layout template and the specific page template
	tmplPath := filepath.Join("templates", tmpl+".html")
	layoutPath := filepath.Join("templates", "layout.html")
	t, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the combined template
	err = t.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
