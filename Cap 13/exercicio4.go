package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) mostrapessoa () {
	fmt.Println(p.nome, p.sobrenome, "", p.idade)
}


func main(){
	anao := pessoa {
		nome: "Dwarf",
		sobrenome: "Guard",
		idade: 150,}

	anao.mostrapessoa()
}