package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	// Essa variável vai designar todos os templates que iremos usar
	// Nesse caso estamos referenciando todos os arquivos .html que estiverem na pasta

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{
			"Murilo",
			"Murilo.klug@gmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
