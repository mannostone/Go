package main

import "fmt"

func main() {
	ss := [][]string{
		[]string {"nome", "sobrenome", "hobby"},
		[]string {"manno", "stone", "games"},
		[]string {"Claudio", "Trindade", "Filmes"},
	}	

	for _, v := range ss{
		for _, item := range v{
			fmt.Printf("%v\t", item)
		}
		fmt.Println()
	}
}
