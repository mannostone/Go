package main

import "fmt"

func main() {
	x := 10
	if x < 100 {
		fmt.Println("100 é maior")
	} else if (x > 100) {
		fmt.Printf("%v é menor que 100", x)
	} else {
		fmt.Printf("%v igual a 100", x)
	}
}
