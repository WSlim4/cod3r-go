package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("notaFinal", notaFinal)

	//cuidado...
	fmt.Println("Teste " + string(123)) // converte o inteiro para o representante na tabela ASCI e não a string literal

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")

	fmt.Println(num - 122)

	//boolean para string

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}
}
