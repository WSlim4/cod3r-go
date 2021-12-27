package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Abacaxi",
		preco:    10.50,
		desconto: 5.50,
	}
	fmt.Println(produto1)
	produto2 := produto{"Salsicha", 1.79, 0.05}
	fmt.Println(produto2.precoComDesconto())
}
