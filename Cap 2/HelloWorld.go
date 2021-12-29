package main

import "fmt"

func main() {
	// Entre dois blocos de texto sempre haverá espaço
	fmt.Println("Olá mundo, gogogo!","Como você está?");

	// Contabilizar numero de bytes e erros
	numerodebytes, erros := fmt.Println("Olá mundo, gogogo!","Como você está?");
	fmt.Println(numerodebytes, erros)

	// Optar por não ter um dos retornos com "_"
	numerodebytes1, _:= fmt.Println("Olá mundo, gogogo!","Como você está?");
	fmt.Println(numerodebytes1)
}
