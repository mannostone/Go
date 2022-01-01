package main

import "fmt"

func main() {
	// []int,tamanho,capacidade
	slice := make([]int,5,10)

	slice[0], slice[1], slice[2], slice[3] = 1,2,3,4
	
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	
	fmt.Println(slice, len(slice), cap(slice))
	// make vai acrescer as "gavetas do slice" aumentando a capacidade
	slice = append(slice, 10)

	fmt.Println(slice, "tamanho: ",len(slice),"capacidade: ", cap(slice))

}
