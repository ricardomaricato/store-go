package routes

import (
	"net/http"

	"github.com/ricardomaricato/store-go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
