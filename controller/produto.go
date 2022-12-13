package controller

import (
	"html/template"
	"loja-go/model"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos, err := model.All()
	if err != nil {
		panic(err.Error())
	}

	temp.ExecuteTemplate(w, "index", produtos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "create", nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		precoConvertido, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			panic(err.Error())
		}

		produto := model.Produto{
			Nome:       r.FormValue("nome"),
			Descricao:  r.FormValue("descricao"),
			Preco:      precoConvertido,
			Quantidade: quantidadeConvertida,
		}

		model.Store(produto)
	}

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produto_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	produto, err := model.Get(produto_id)

	if err != nil {
		panic(err.Error)
	}

	temp.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto_id, _ := strconv.Atoi(r.FormValue("id"))
		produto, err := model.Get(produto_id)

		if err != nil {
			panic(err.Error)
		}

		precoConvertido, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			panic(err.Error())
		}

		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		produto.Preco = precoConvertido
		produto.Quantidade = quantidadeConvertida

		model.Update(produto)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	produto_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := model.Delete(produto_id)

	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", 301)
}
