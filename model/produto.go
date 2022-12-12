package model

import (
	"fmt"
	"loja-go/config"
)

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

func InsertProduto(produto Produto) error {
	db := config.GetConexao()

	_, err := db.Exec(fmt.Sprintf("INSERT INTO produto (nome, descricao, preco, quantidade) VALUES ('%s','%s',%f,%d)", produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	if err != nil {
		return err
	}

	fmt.Println("Produto cadastrado com sucesso.")

	return nil
}

func GetProdutos() ([]*Produto, error) {
	db := config.GetConexao()

	res, err := db.Query("SELECT * FROM produto")

	if err != nil {
		return nil, err
	}

	produtos := []*Produto{}
	for res.Next() {
		var produto Produto
		if err := res.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade); err != nil {
			return nil, err
		}

		produtos = append(produtos, &produto)
	}

	return produtos, nil
}
