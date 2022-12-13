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

func All() ([]*Produto, error) {
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

func Get(id int)(*Produto, error){
	db := config.GetConexao()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM produto WHERE id = %d", id))

	if err != nil {
		return nil, err
	}

	var produto Produto
	for res.Next(){
		if err := res.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade); err != nil {
			return nil, err
		}
	}

	return &produto, nil
}

func Store(produto Produto) error {
	db := config.GetConexao()

	_, err := db.Exec(fmt.Sprintf("INSERT INTO produto (nome, descricao, preco, quantidade) VALUES ('%s','%s',%f,%d)", produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	if err != nil {
		return err
	}

	fmt.Println("Produto cadastrado com sucesso.")

	return nil
}

func Update(produto *Produto) error{
	db := config.GetConexao()

	_, err := db.Exec(fmt.Sprintf("UPDATE produto SET nome = '%s', descricao = '%s', preco = %f, quantidade = %d WHERE id = %d", produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id))

	fmt.Println(err)
	if err != nil {
		return err
	}

	fmt.Println("Produto alterado com sucesso.")

	return nil
}

func Delete(produto_id int) error {
	db := config.GetConexao()

	_, err := db.Exec(fmt.Sprintf("DELETE FROM produto WHERE id = %d", produto_id))

	if err != nil {
		return err
	}

	fmt.Println("Produto excluído com sucesso")

	return nil
}
