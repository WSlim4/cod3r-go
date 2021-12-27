package main

import "fmt"

func main() {
	// homogêneo (mesmo tipo) e estático (tamanho fixo)

	var notas [3]float64

	notas[0], notas[1], notas[2] = 7.4, 6.3, 5.1

	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f\n", media)
}
