package main

import "fmt"



func main(){
	x := 189
	//função anonima
	//não possuem nome, apenas são criadas e chamadas em seguida
	func (x int) {
		fmt.Println(x, "x 545441 =")
		fmt.Println(x*545441)
	}(x)//chamando a função anonima


}