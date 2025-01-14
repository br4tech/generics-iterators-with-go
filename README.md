# Generics em Go, a forma de lidar com coleções passou por uma revolução!

Agora, com a ajuda dos iterators, podemos repensar como interagimos com conjuntos de dados, abrindo um leque de possibilidades para um código mais limpo, eficiente e reutilizável.

Obs: Podemos utilizar generic apartir da versao 1.18 do Go 🤓  

## O problema das coleções antes dos Generics

- **Código repetitivo**: Imagine ter que escrever funções separadas para processar slices de `int`, `string`, `float64`, etc. Era um festival de `for` loops com lógica praticamente idêntica, mudando apenas o tipo de dado. 😪  
- **Funções menos flexíveis**: Sem Generics, era difícil criar funções que operassem em diferentes tipos de coleções. Imagine uma função para filtrar elementos de um slice - você precisaria de uma versão para cada tipo de dado! 🤯  
- **Dificuldade em encapsular lógica**: Reutilizar a lógica de iteração em diferentes partes do código era um desafio, levando à duplicação e dificultando a manutenção. 😫  

## A solução com Generics e Iterators

- **Abstração e reutilização**: Com Generics, podemos escrever funções e estruturas de dados que operam em tipos arbitrários, sem precisar de código repetitivo. Crie uma função `filter` que funciona com qualquer tipo de slice! 🤩  
- **Iterators para simplificar a iteração**: Os iterators fornecem uma maneira padronizada de percorrer coleções, independentemente de sua estrutura interna. Imagine um `for...range` que funciona com qualquer tipo de coleção, de slices a mapas e até estruturas de dados personalizadas! 🤯  
- **Encapsulamento de lógica**: Abstraia a lógica de iteração em iterators, permitindo que você a reutilize em diferentes partes do código e simplificando a manutenção. ✨  
- **Código mais limpo e legível**: Com menos repetição e maior abstração, seu código se torna mais conciso, fácil de entender e manter. 🤓  

## Exemplos práticos

1. **Criando um iterator para um tipo personalizado**  
   Imagine uma estrutura de dados `ListaEncadeada`. Você pode implementar um iterator para ela e usar `for...range` para percorrer seus elementos! 🤯  

2. **Filtrando elementos de qualquer slice**  
   Crie uma função `filter` genérica que recebe um slice de qualquer tipo e uma função de predicado. Ela retorna um novo slice contendo apenas os elementos que satisfazem o predicado. 🤩  

3. **Implementando um `map` genérico**  
   Crie uma função `map` que recebe um slice de qualquer tipo e uma função de transformação. Ela retorna um novo slice com os elementos transformados. 🚀  

## Conclusão

Generics e iterators em Go permitem que você escreva código mais flexível, reutilizável e legível ao lidar com coleções. Explore essas ferramentas para criar soluções mais elegantes e eficientes! 💪  

## Executar

Para executar esse script podemos rodar em nosso terminal o comando:

```bash
go run cmd/main.go
```

Devemos ter como resultado, uma entidade de produto pode ser um livro e um eletronico:

```bash
{O Senhor dos Anéis 79.9 10 {J.R.R. Tolkien 978-8533629437}}
{Smartphone Galaxy S23 3999.9 5 {Samsung S23}}
```

## Quer saber mais?

- [Documentação do Go](https://go.dev/doc/)  
- [Tutorial de Generics](https://go.dev/doc/tutorial/generics)  
- [Exemplo de Iterator](https://refactoring.guru/pt-br/design-patterns/iterator/go/example)  

#golang #generics #iterators #coleções #programação
