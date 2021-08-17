package main

import (
	"html/template"
	"net/http"

	"github.com/store-go/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)

}
