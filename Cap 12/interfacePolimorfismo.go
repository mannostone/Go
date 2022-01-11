package main

import "fmt"

type gente struct {
	nome string
	idade int
}

type dentista struct {
	gente
	dentesarrancados int
	salario float64

}
type arquiteto struct {
	gente
	tipodeconstrucao string
	tamanhodaloucura string
}

func (x dentista) oibomdia(){
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}
func (x arquiteto) oibomdia(){
	fmt.Println("Meu nome é", x.nome, "e ouve só: Hoje estou fazendo ", x.tipodeconstrucao )
}

type pessoa interface {
	oibomdia()
}

func serhumanno(p pessoa) {
	p.oibomdia()
}

func main(){
	//dentista
	homemdente := dentista {
		// atributos de "type gente"
		gente: gente{
		nome : "ArrancaTudo",
		idade: 36,
	},
	// atributos do método
	dentesarrancados: 975,
	salario: 19438}
	//arquiteto
	predioman := arquiteto {
		// atributos de "type gente"
		gente: gente{
			nome: "Beatriz Toledo",
			idade: 22,
		},
		// atributos do método
		tipodeconstrucao: "Reconstrução de Brumadinho",
		tamanhodaloucura: "Doida",}
		
		// forma de chamar uma função
		homemdente.oibomdia()
		predioman.oibomdia()

		// forma alternativa de chamar 
		serhumanno(predioman)
		serhumanno(homemdente)
}