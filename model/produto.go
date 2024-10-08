package model

import (
	"Lojinha-DigiPort/loja-digiport-backend/db"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type Produto struct {
	Id                  string  `json:"id"`
	Nome                string  `json:"nome"`
	Preco               float64 `json:"preco"`
	Descricao           string  `json:"descricao"`
	Imagem              string  `json:"imagem"`
	QuantidadeEmEstoque int     `json:"quantidade"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")
	// err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)

	if err != nil {
		panic(err.Error())

	}

	p := Produto{}
	produtos := []Produto{}

	for resultado.Next() {

		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()
	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto nao encontrado %s\n", nome)
	} else if err != nil {
		panic(err.Error())
	}

	var produto1 Produto
	produto1.Id = id
	produto1.Nome = nome
	produto1.Preco = preco
	produto1.Descricao = descricao
	produto1.Imagem = imagem
	produto1.QuantidadeEmEstoque = quantidade

	defer db.Close()
	return produto1

}

func CriaProduto(prod Produto) error {

	if produtoCadastrado(prod.Nome) {
		fmt.Printf("Produto ja cadastrado: %s\n", prod.Nome)
		return fmt.Errorf("Produto ja cadastrado")
	}

	db := db.ConectaBancoDados()
	id := uuid.NewString()
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	strInsert := "INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)"

	result, err := db.Exec(strInsert, id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade))
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Produto %s criado com sucesso ( %d row affected)\n", id, rowsAffected)

	defer db.Close()

	return nil

}

func produtoCadastrado(nomeProduto string) bool {
	prod := BuscaProdutoPorNome(nomeProduto)
	return prod.Nome == nomeProduto

}

func RemoveProduto(id string) error {
	db := db.ConectaBancoDados()
	defer db.Close()

	//DELETE FROM [nome da tabela] WHERE [nome do campo] = [valor]
	resultado, err := db.Exec("DELETE FROM PRODUTOS WHERE id = $1", id)
	if err != nil {
		fmt.Println("erro ao tentar excluir produto")
		return fmt.Errorf("ocorreu um erro ao tentar excluir produto: %w", err)
	}

	linesAffected, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("%d linhas afetadas", linesAffected)

	fmt.Println("produto excluido")

	return nil

}
