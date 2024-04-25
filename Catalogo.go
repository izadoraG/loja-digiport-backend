package main

import (
	"Lojinha-DigiPort/loja-digiport-backend/model"
)

func ListaProdutos() []model.Produto {
	produtos := []model.Produto{
		{Nome: "Blusa Floral", Descrição: "Blusa de flores estampadas", Categoria: "Roupas", Id: "001", Valor: 29.99, Quantidade: 10, Imagem: "imagem_blusa_floral.jpg"},
		{Nome: "Calça Jeans", Descrição: "Calça jeans azul", Categoria: "Roupas", Id: "002", Valor: 39.99, Quantidade: 5, Imagem: "imagem_calca_jeans.jpg"},
		{Nome: "Vestido Vintage", Descrição: "Vestido vintage estampado", Categoria: "Roupas", Id: "003", Valor: 49.99, Quantidade: 3, Imagem: "imagem_vestido_vintage.jpg"},
		{Nome: "Bolsa de Couro", Descrição: "Bolsa de couro marrom", Categoria: "Acessórios", Id: "004", Valor: 59.99, Quantidade: 8, Imagem: "imagem_bolsa_couro.jpg"},
	}
	return produtos
}

func BuscaProdutoPorNome(nome string) []model.Produto {
	var produtosEncontradosporNome []model.Produto // variavel que armazena produtos encontrados

	for _, produto := range ListaProdutos() {
		if produto.Nome == nome {
			produtosEncontradosporNome = append(produtosEncontradosporNome, produto)
		}
	}
	return produtosEncontradosporNome

}
