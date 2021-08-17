package main

import (
	"net/http"

	"github.com/store-go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
