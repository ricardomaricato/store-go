package routes

import (
	"net/http"

	"github.com/ricardomaricato/store-go/controllers"
)

//comentário

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
