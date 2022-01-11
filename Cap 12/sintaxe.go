package main

import "fmt"

func main(){
	valor := soma(10,10)

	fmt.Println(valor)
}
//   nome(assinatura e tipo) retorno {}
func soma(x int , y int) int {
	return x+y
}