package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"));

type Produto struct{
	Nome, Descricao string
	Preco float64
	Quantidade int
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos := []Produto{
		{"Camiseta","Confortável",180,15},
		{"Tênis","12 molas",180,15},
		{"Boné","Boné lagum",180,15},
	}
	
	temp.ExecuteTemplate(w, "index", produtos)
}