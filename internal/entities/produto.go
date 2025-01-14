package entities

type Produto[T any] struct {
	Nome      string
	Preco     float64
	Estoque   int
	Atributos T
}
