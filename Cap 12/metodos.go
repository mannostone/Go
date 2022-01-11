package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}
// método específico para um "tipo"
func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main(){
	mauricio := pessoa{"Mauricio",30}
	mauricio.oibomdia()
}