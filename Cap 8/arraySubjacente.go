package main

import "fmt"

func main() {
	primeiroslice := []int{1,2,3,4,5}
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)
	// fatia do primeiro slice
	fmt.Println(segundoslice)
	// primeiroslice modificado (?)
	fmt.Println(primeiroslice)
}
