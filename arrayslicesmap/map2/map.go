package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Emuxo":    12342.05,
		"Lanzinho": 15477.00,
		"Abinho":   1000.45,
	}

	funcsESalarios["Wesley"] = 2500.00
	delete(funcsESalarios, "Pancracio")

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s salario: %.2f\n", nome, salario)
	}

}
