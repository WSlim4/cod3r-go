package main

import "fmt"

func main() {
	// ambos os slices apontam para o mesmo array interno
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	s1[0] = 2

	fmt.Println(s1, s2)

}
