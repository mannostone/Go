package main

import "fmt"

func main() {
	x := 1
	switch {
	case x == 0:
		fmt.Println("Zero")
	case x == 1:
		fmt.Println("Um")
	case x == 2:
		fmt.Println("Dois")

	}
}
