package controllers

import (
	"html/template"
	"net/http"

	"github.com/ricardomaricato/store-go/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	template.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {
	template.ExecuteTemplate(w, "New", nil)
}
