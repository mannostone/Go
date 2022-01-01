package main

import "fmt"

func main() {
	mepe := map[string][]string{
		"Stone,Manno": []string{"Games","Indies"},
		"Trindade,Claudio": []string{"Filmes", "Séries"},
		"Toledo,Beatriz": []string{"Fornite"}, // kkk
	}

	// Mostra o nome da pessoa (k)
	// v são os "hobbies"
	for k,v := range mepe {
		fmt.Println(k)
		for i, h := range v { // range v são a quantidade de atributos
			fmt.Println("\t",i,h) // i é o indice e h o valor do indice atual
		}
	} 
}
