package main

import (
	"database/sql"
	"html/template"
	"loja-go/config"
	"loja-go/model"
	"net/http"
)

var (
	db *sql.DB
	err error
	temp = template.Must(template.ParseGlob("templates/*.html"));
)

func main(){
	db = config.GetConexao()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos, err := model.GetProdutos()
	if err != nil {
		panic(err.Error())
	}
	
	temp.ExecuteTemplate(w, "index", produtos)
}