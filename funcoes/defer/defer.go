package main

import "fmt"

func obterValorAprovado(n int) int {
	defer fmt.Println("Shipancoli!")
	if n > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return n

}

func main() {
	fmt.Println(obterValorAprovado(5001))
}
