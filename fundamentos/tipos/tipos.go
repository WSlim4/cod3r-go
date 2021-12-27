package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode(int32)
	fmt.Println("O tipo do inteiro é", reflect.TypeOf(i2))
	fmt.Println("O valor do rune é ", i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99))

	//boolean
	var bo bool = true
	fmt.Println("O tipo da variável bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string

	str := "mundo"
	fmt.Println("Olá " + str)
	fmt.Println("O tamanho da string é", len(str))

	str2 := `Olá
	meu
	nome
	é
	Wesley`

	fmt.Println("O tamanho da string é", len(str2))
}
