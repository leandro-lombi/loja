package main

import (
	"net/http"
	"text/template"
)

// Essa função Must encapsula todos os templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

func main()  {
	// Toda requisição para o /, quem irá atender será o index
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	templates.ExecuteTemplate(w, "Index", nil)
}