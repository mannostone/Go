package main

import "fmt"

func main(){
	// defer = executa por último
	// caso tenha vários, ele aplica o conceito de pilha
	defer fmt.Println("com defer")
	fmt.Println("sem defer")
}