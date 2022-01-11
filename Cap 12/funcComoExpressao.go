package main

import "fmt"

func main(){
	x := 189
	// função como expressão/valor
	y := func (x int) int {
		// fmt.Println(x, "x 545441 =")
		return x*545441
	}
	//
	fmt.Println(x, "x 545441 =", y(x))
}