package main

import "fmt"

func main() {
	x := 0
	// estarecebeovalor(x)
	// estarecebeumponteiro(&x)
	
	// estarecebeovalor(x) = 0
	// estarecebeumponteiro(x) = 1
	fmt.Println("no main tem: ", x)

}

func estarecebeovalor(x int) {
	// copia do valor passado por argumento
	x++
	fmt.Println("na função tem: ", x)

}
func estarecebeumponteiro(x *int) {
	// endereço passado por argumento
	*x++
	fmt.Println("Na função tem: ", *x)

}
