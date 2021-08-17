package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 3},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	tmpl.ExecuteTemplate(w, "Index", produtos)
}
