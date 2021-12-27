package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678] = "Pancracio"
	aprovados[13246795] = "Emuxo"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 12345678)
	fmt.Println(aprovados)
}
