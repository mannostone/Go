package main

import "fmt"

var x int = 10

func main() {
	// erro -> x = 10.5
	fmt.Printf("%t", x)
}