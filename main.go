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
	Id, Quantidade int
	Nome, Descricao string
	Preco float64
}

func insertProduto(produto Produto) error{
	_, err := db.Exec(fmt.Sprintf("INSERT INTO produto (nome, descricao, preco, quantidade) VALUES ('%s','%s',%f,%d)", produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	if(err != nil){
		return err
	}

	fmt.Println("Produto cadastrado com sucesso.")

	return nil
}

func getProdutos() ([]*Produto, error){
	res, err := db.Query("SELECT * FROM produto")

	if err != nil {
		return nil, err
	}

	produtos := []*Produto{}
	for res.Next(){
		var produto Produto
		if err := res.Scan(&produto.Id ,&produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade); err != nil{
			return nil, err
		}

		produtos = append(produtos, &produto)
	}

	return produtos, nil
}

func main(){

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/loja_go")

	produto := Produto{
		4,
		5,
		"Bermuda",
		"Da cyclone",
		89.99,
	}

	if inserterror := insertProduto(produto); inserterror != nil {
		fmt.Println(inserterror)
		panic(err)
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos, err := getProdutos()
	if err != nil {
		panic(err.Error())
	}
	
	temp.ExecuteTemplate(w, "index", produtos)
}