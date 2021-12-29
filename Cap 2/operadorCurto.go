package main

import "fmt"
// Variáveis visiveis para todo o pacote (Package Level Scope)
var HK = "speedrun 40:24.52"

func main(){
// declarador de variáveis :=
	x := 10
	fmt.Println(x)

	y := 20
	// %v valor / %t tipo
	fmt.Printf("y: %v, %T", y, y)

	x = y
	fmt.Printf("\nx: %v, %T\n", x, x)

	fmt.Println(HK)
}