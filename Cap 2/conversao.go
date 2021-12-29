package main

import "fmt"

type hotdog int
var b hotdog

func main(){
	x := 10
	fmt.Printf("%t\n", x)
	fmt.Printf("%t\n", b)

	// conversão
	x = int(b)
	fmt.Printf("%t", x)


}