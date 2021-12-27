package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	fmt.Println(numeros)

	for i, numero := range numeros {
		fmt.Println("Elemento", numero)
		fmt.Println("Índice", i)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
