package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415

	var raio = 3.2 // inferido o tipo ao atribuir o valor a variável

	area := PI * m.Pow(raio, 2)

	fmt.Println("Area da circunferência", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := false, 2, "Oi"

	fmt.Println(g, h, i)
}
