package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(6))
	fmt.Println(loops(6))

}

// recursividade
func fatorial(x int) int {
	if x == 1 {
		return x
	}
	// a função chama ela mesma até que seja 1
	return x * fatorial(x-1)
}
//recursividade com loop (mais seguro)
func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
