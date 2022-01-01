package main

import "fmt"

func main() {
	x := make([]string,26,26)
	slice := []string{ "Acre", "Alagoas", 
	"Amapá", "Amazonas", "Bahia", "Ceará", 
	"Espírito Santo", "Goiás", "Maranhão", 
	"Mato Grosso", "Mato Grosso do Sul", 
	"Minas Gerais", "Pará", "Paraíba", 
	"Paraná", "Pernambuco", "Piauí", 
	"Rio de Janeiro", "Rio Grande do Norte", 
	"Rio Grande do Sul", "Rondônia", "Roraima", 
	"Santa Catarina", "São Paulo", "Sergipe", 
	"Tocantins"}

	x = slice

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}
