package routes

import (
	"net/http"

	"github.com/ricardomaricato/store-go/controllers"
)

//comentario

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
