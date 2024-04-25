package main

import (
	"Lojinha-DigiPort/loja-digiport-backend/model"
	"errors"
)

var Produtos []model.Produto = []model.Produto{}

func criarcatalogo(produtoNovo model.Produto) error {
	for _, produto := range Produtos {
		if produtoNovo.Id == produto.Id {
			return errors.New("produto com Esse ID ja existente")

		}
	}
	Produtos = append(Produtos, produtoNovo)
}

func BuscaProdutoPorNome(nome string) []model.Produto {
	produtosEncontradosporNome := []model.Produto{} // variavel que armazena produtos encontrados

	for _, produto := range Produtos {
		if produto.Nome == nome {
			produtosEncontradosporNome = append(produtosEncontradosporNome, produto)
		}
	}
	return produtosEncontradosporNome

}
