package main
import "fmt"

type pessoa struct {
	nome string
	idade int
}

type profissional struct {
	pessoa // struct declarada acima
	titulo string
	salario int
}

func main(){
	pessoa1 := pessoa {"Alfredo", 30}
	fmt.Println()
	/* pessoa2 := profissional {
		pessoa {
			nome = "Juliana",
			idade = 25,
		},
		titulo = "aaa",
		salario = 2000} */
	pessoa2 := profissional {pessoa {"Juliana", 25},"aaa",2000}
	// mostrar um atributo específico
	// fmt.Println(pessoa1.idade)
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	
}