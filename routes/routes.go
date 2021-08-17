package routes

import (
	"net/http"

	"github.com/store-go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
