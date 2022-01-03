package main
import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sabores []string
}


func main(){
	meumapa := make(map[string]pessoa)
	// criando as pessoas
	pessoa1 := pessoa {
		nome: "Alfredo", 
		sobrenome: "Wagner",
		sabores: []string{"Chocolate", "morango", "creme"}}
	// forma alternativa de criar
	pessoa2 := pessoa {"Juliana","Borges", 
	[]string{"Creme","purpurina","napolitano"}}
	
	// adicionando essas pessoas ao mapa (pode ser adicionado diretamente como em "Cap 8/maps.go")
	meumapa [pessoa1.sobrenome] = pessoa1
	meumapa [pessoa2.sobrenome] = pessoa2


	for _, v := range meumapa{
		fmt.Println(v.sobrenome, v.nome)
		for _, dados := range v.sabores {
			fmt.Println("->", dados)

		}
	} 
	
}