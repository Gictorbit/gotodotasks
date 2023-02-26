package main

import (
	"html/template"
	"net/http"
)

func NewWebServer() http.Handler {
	mux := http.NewServeMux()
	// Define the route handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	mux.HandleFunc("/login", showLogin)
	mux.HandleFunc("/taskmanager", showTaskManager)

	// Serve the static files
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets/"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css/"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js/"))))

	return mux
}

func showLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Domain string
	}{
		Domain: Domain,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func showTaskManager(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/taskmanager.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Domain string
	}{
		Domain: Domain,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
