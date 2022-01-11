package main

import "fmt"

func main() {
	x := 0
	//y = endereço de x
	y := &x
	// valor no qual y aponta
	fmt.Println(*y)

	*y++
	fmt.Println(x, *y)
	fmt.Printf("%T %T", x, y)
}
