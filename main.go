package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// Route pour robots.txt
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "User-agent: *")
		fmt.Fprintln(w, "Disallow: /public/")
	})

	// Routes principales
	http.HandleFunc("/", handler)
	http.HandleFunc("/projet1", handler)
	http.HandleFunc("/projet2", handler)
	http.HandleFunc("/projet3", handler)
	http.HandleFunc("/bomberman", handler)

	// Fichiers publics
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// Démarrage du serveur
	fmt.Println("Server started on port 8080 : http://localhost:8080")
	logfile("\n______________________________________________________\n___________________Server started_____________________\n______________________________________________________\n")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var adress = ""
	switch r.URL.Path {
	case "/":
		adress = "./public/templates/index.html"
	case "/projet1":
		adress = "./public/templates/projet1.html"
	case "/projet2":
		adress = "./public/templates/projet2.html"
	case "/bomberman":
		adress = "./public/templates/bomberman.html"
	case "/projet3":
		adress = "./public/templates/projet3.html"
	default:
		http.NotFound(w, r)
		return
	}

	logfile(r.URL.Path)

	ts, err := template.ParseFiles(adress)
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		log.Printf("Error loading template: %v\n", err)
		return
	}
	ts.Execute(w, nil)
}

func logfile(texttowrite string) {
	// Open or create the log file
	logFile, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v\n", err)
	}

	// Set the logger to write to the file
	log.SetOutput(logFile)

	// Log only the path accessed for statistics
	log.Printf("Page : %s\n", texttowrite)

	logFile.Close()
}
