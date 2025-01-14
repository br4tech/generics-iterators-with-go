package service

import "github.com/br4tech/generics-iterators-with-go/internal/entities"

type ListaProdutos[T any] struct {
	Produtos []entities.Produto[T]
}

func (l *ListaProdutos[T]) AdicionarProduto(p entities.Produto[T]) {
	l.Produtos = append(l.Produtos, p)
}

func (l *ListaProdutos[T]) Iterador() <-chan entities.Produto[T] {
	ch := make(chan entities.Produto[T])
	go func() {
		for _, p := range l.Produtos {
			ch <- p
		}
		close(ch)
	}()
	return ch
}
