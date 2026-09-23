package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/home.html
var templateFiles embed.FS

func main() {
	// Parse the template once when the server starts instead of on every request.
	tmpl, err := template.ParseFS(templateFiles, "templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	// Keep the root URL friendly by sending visitors to the home page.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		homeHandler(tmpl, w, r)
	})
	fmt.Println("starting the server on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(tmpl *template.Template, w http.ResponseWriter, r *http.Request) {
	// This handler owns only the /home page; unknown paths should return 404.
	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	// The page only supports GET requests because it does not accept form data.
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Set the content type before writing the rendered HTML response.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		log.Printf("render home template: %v", err)
	}
}
