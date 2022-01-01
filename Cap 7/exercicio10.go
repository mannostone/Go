package main

import "fmt"

func main() {
	// true
	fmt.Println(true && true)
	// false
	fmt.Println(true && false)
	// true
	fmt.Println(true || true)
	// true
	fmt.Println(true || true)
	// false
	fmt.Println(!true)
}
