package main

import "fmt"

func main() {
	//recebe o retorno da função
	x := retornaumafuncao()
	x()
	
}
//   método             anônimo   retorno
func retornaumafuncao() func() {
	return func(){
		fmt.Println("Voltei")
	}
}
