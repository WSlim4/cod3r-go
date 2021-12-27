package main

import "fmt"

func main() {
	// Go não tem aritmética de ponteiros

	i := 1

	var p *int = nil

	p = &i // Pegando o endereço de memória de i e atribuindo ao ponteiro p
	*p++   // Desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
