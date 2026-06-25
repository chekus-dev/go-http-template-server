package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", homeHandler)
	fmt.Println("starting the server on port :7000")
	if err := http.ListenAndServe(":7000", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) { // Haandles request to /home
	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

}
