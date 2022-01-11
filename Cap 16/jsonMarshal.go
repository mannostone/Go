package main
import "encoding/json"
import "fmt"
// json só faz importações de métodos ou seja o que for se estiverem com a 1° letra maiúscula
type Pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	Profissao string
	ContaBancaria float64
}

func main(){
	jamesBond := Pessoa{
		Nome : "James",
		Sobrenome: "Bond",
		Idade: 40,
		Profissao: "Agente não tão secreto",
		ContaBancaria: 8001.01,
	}

	mS := Pessoa {"Manno","Stone", 23,"Desenvolvedor?",0.06}

	j, err := json.Marshal(jamesBond)
	if err != nil{fmt.Println(err)}
	d, err := json.Marshal(mS)
	if err != nil{fmt.Println(err)}

	fmt.Printf(string(j),"\n")
	fmt.Printf(string(d))


}