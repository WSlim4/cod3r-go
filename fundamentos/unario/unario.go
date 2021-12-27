package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Go possui apenas o incremento pos fixado, ou seja, o operador unário vem após a variável

	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)
}
