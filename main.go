package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Chargement de la page principale
	http.HandleFunc("/", HomeHandler)

	// Fichiers annexes
	http.Handle("/public/static/", http.StripPrefix("/public/static/", http.FileServer(http.Dir("./public/static/"))))
	http.Handle("/public/js/", http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))
	http.Handle("/public/img/", http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))

	fmt.Println("Server started on port 8080 : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./public/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	ts.Execute(w, nil)
}
