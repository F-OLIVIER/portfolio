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
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/static/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./public/img/"))))

	fmt.Println("Server started on port 8080 : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	ts.Execute(w, nil)
}
