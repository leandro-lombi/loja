package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Essa função Must encapsula todos os templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Toda requisição para o /, quem irá atender será o index
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// slice de produto
	produtos := []Produto{
		{"Camiseta", "Algodão", 29.9, 2},
		{"Tenis", "Azul", 349.5, 4},
		{"Meia", "Branca", 8.49, 15},
		{"Sapato", "Marrom", 143.75, 6},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
