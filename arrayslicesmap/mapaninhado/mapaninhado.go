package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"P": {
			"pancracio":    15487.04,
			"pancreitinho": 14572.75,
		},
		"E": {
			"emuxo": 12487.74,
		},
	}

	delete(funcsPorLetra, "E")

	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
