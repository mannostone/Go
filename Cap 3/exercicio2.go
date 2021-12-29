package main

import "fmt"

var x int
var y string
var z bool

func main(){
	fmt.Printf("%v /t %v /t %v", x,y,z)
	// Variáveis criadas e não inicializadas possuem valor "0" diferente
}