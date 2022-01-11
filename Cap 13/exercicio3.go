package main

import "fmt"

func main(){
	si := []int {10,10,1,2,3,4,5,6,8,9,23}
	// forma de passar slices como argumentos
	defer fmt.Println("com defer", soma(si...))
	fmt.Println("sem defer", soma2(si))
}
//   nome(assinatura variádica) retorno {}
func soma(x... int) int {
	soma := 0
	for _, v := range x{
		soma+=v
	}
	return soma
}

func soma2(x []int) int {
	total := 0
	for _, v:= range x {
		total += v
	}
	return total
}