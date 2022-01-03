package main
import "fmt"

type cliente struct {
	nome string
	sobrenome string
	fumante bool
}

func main(){
	// forma de declarar
	cliente1 := cliente {
		nome: "João",
		sobrenome: "Ávila",
		fumante: false,
	}
	// forma de declarar
	cliente2 := cliente {"Joana","Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}