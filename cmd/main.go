package main

import (
	"fmt"

	"github.com/br4tech/generics-iterators-with-go/internal/entities"
	"github.com/br4tech/generics-iterators-with-go/internal/service"
)

func main() {

	livro := entities.Produto[interface{}]{
		Nome:    "O Senhor dos Anéis",
		Preco:   79.90,
		Estoque: 10,
		Atributos: entities.Livro{
			Autor: "J.R.R. Tolkien",
			ISBN:  "978-8533629437",
		},
	}

	eletronico := entities.Produto[interface{}]{
		Nome:    "Smartphone Galaxy S23",
		Preco:   3999.90,
		Estoque: 5,
		Atributos: entities.Eletronico{
			Marca:  "Samsung",
			Modelo: "S23",
		},
	}

	lista := service.ListaProdutos[interface{}]{} // Inicializa a lista
	lista.AdicionarProduto(livro)                 // Adiciona os produtos usando o método
	lista.AdicionarProduto(eletronico)

	for produto := range lista.Iterador() {
		fmt.Println(produto)
	}
}
