package main

import (
	"fmt"
)

func main() {
	// chama a função somentePares e passa os valores
	t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	// vai armazenar apenas os inteiros
	var slice []int
	for _, v := range y { // range que vai percorrer
		if v%2 == 0 {
			//slice
			slice = append(slice, v)
		}
	}
	//"callback"
	// total recebe a soma dos valores pares
	// "f(slice...) é a função soma"
	total := f(slice...)
	// valor exibido em main
	return total 
}
