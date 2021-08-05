package routes

import (
	"net/http"

	"github.com/lnl/loja/controllers"
)

func CarregaRotas() {
	// Toda requisição para o /, quem irá atender será o index
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
