package main

import "fmt"

func main() {
	esporteFavorito := "Corrida"
	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Zero")
	case "Volei":
		fmt.Println("Um")
	case "Hoquei":
		fmt.Println("Dois")
	case "Corrida":
		fmt.Println("Hamilton perdeu na ultima volta")

	}
}
