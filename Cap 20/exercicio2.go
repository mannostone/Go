package main
import "fmt"

type pessoa struct {
	nome string
	idade int
}

func (p *pessoa) falar () {
	fmt.Println(p.nome, "disse oi")
}
// toda struct que tenha os métodos abaixos será incluida na interface humanos 
type humanos interface {
	falar()
}

func dizerAlgumaCoisa (h humanos) {
	h.falar()
}

func main(){
	p1 := pessoa {"Manno Stone", 23}

	// ou (&p1).falar()
	p1.falar()
	// método que pertence a classe humanos
	dizerAlgumaCoisa(&p1)

}