package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
	temp = template.Must(template.ParseGlob("templates/*.html"));
)


type Produto struct{
	Nome, Descricao string
	Preco float64
	Quantidade int
}

func insertProduto(produto Produto) error{
	fmt.Println("Chega até aqui")
	_, err := db.Exec(fmt.Sprintf("INSERT INTO produto (nome, descricao, preco, quantidade) VALUES ('%s','%s',%f,%d)", produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	if(err != nil){
		return err
	}

	fmt.Println("Produto cadastrado com sucesso.")

	return nil
}

func main(){

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/loja_go")

	produto := Produto{
		"Camiseta",
		"Confortável",
		89.99,
		15,
	}

	if inserterror := insertProduto(produto); inserterror != nil {
		fmt.Println(inserterror)
		panic(err)
	}

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