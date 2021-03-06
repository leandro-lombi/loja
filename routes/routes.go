package routes

import (
	"net/http"

	"github.com/lnl/loja/controllers"
)

func CarregaRotas() {
	// Toda requisição para o /, quem irá atender será o index
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
