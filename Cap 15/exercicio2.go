package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func mudeMe (p *pessoa) {
	(*p).nome = "Manninho"
	p.sobrenome = "Stoninho"
}


func main() {
	ze := pessoa{"Zé","Polvilho",32}
	fmt.Println(ze)

	mudeMe(&ze)
	fmt.Println(ze)
}