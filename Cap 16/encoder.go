package main
import "encoding/json"
import "os"

type pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	Profissao string
	ContaBancaria float64
}

func main(){
	jamesBond := pessoa{
		Nome : "James",
		Sobrenome: "Bond",
		Idade: 40,
		Profissao: "Agente não tão secreto",
		ContaBancaria: 8001.01,
	}
	//                         interface
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesBond)
}