package main

import "fmt"

func main() {
	// numero de elementos limitado
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	// numero de elementos ilimitado
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	sliceString := []string{"banana", "maça", "jaca"}
	// indice = ciclo loop / valor = dado em sliceString
	for indice, valor := range sliceString {
		fmt.Println("No indice, ", indice, "temos o valor: ", valor)
	}
	// append adiciona valores a um slice
	sliceString = append(sliceString, "melancia")
	for indice, valor := range sliceString {
		fmt.Println("No indice, ", indice, "temos o valor: ", valor)
	}

	total := 0
	for _, valor := range slice {
		total += valor
	}
	fmt.Println("Total", total)

}
