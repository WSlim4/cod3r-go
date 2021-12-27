package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // também suportado no switch
		fmt.Println("Ganhou!!!") // a variável i só está disponível dentro do bloco if else
	} else {
		fmt.Println("Perdeu!!!")
	}
}
