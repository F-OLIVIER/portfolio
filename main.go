package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/projet1", handler)
	http.HandleFunc("/projet2", handler)
	http.HandleFunc("/projet3", handler)
	http.HandleFunc("/bomberman", handler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	fmt.Println("Server started on port 8080 : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var adress = ""
	switch r.URL.Path {
	case "/projet1":
		adress = "./public/templates/projet1.html"
	case "/projet2":
		adress = "./public/templates/projet2.html"
	case "/bomberman":
		adress = "./public/templates/bomberman.html"
	case "/projet3":
		adress = "./public/templates/projet3.html"
	default:
		adress = "./public/templates/index.html"
	}

	ts, err := template.ParseFiles(adress)
	if err != nil {
		log.Fatal(err)
	}
	ts.Execute(w, nil)
}
