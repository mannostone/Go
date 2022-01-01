package main

import "fmt"

func main() {
	mepe := map[string][]string{
		"Stone,Manno": []string{"Games","Indies"},
		"Trindade,Claudio": []string{"Filmes", "Séries"},
		"Toledo,Beatriz": []string{"Fornite"}, // kkk
	}
	// adicionando valor ao mapa
	mepe["Sparrow,Jack"] = []string{"Prover conteúdo gratuito"}
	//deletando um valor do mapa
	delete(mepe,"Sparrow,Jack")
	
	// Mostra o nome da pessoa (k)
	// v são os "hobbies"
	for k,v := range mepe {
		fmt.Println(k)
		for i, h := range v { // range v são a quantidade de atributos
			fmt.Println("\t",i,h) // i é o indice e h o valor do indice atual
		}
	} 
}
