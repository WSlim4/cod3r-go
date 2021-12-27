package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	n1, n2 := troca(1, 2)

	fmt.Println(n1, n2)
}
