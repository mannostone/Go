package main

import (
	"fmt"
)

func main() {
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	
	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0 // fora do scope
	return func() int{
		x++ // usado dentro do scope
		return x
	}
}