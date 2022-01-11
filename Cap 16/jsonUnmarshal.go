package main
import "fmt"
import "encoding/json"
//{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente não tão secreto","ContaBancaria":8001.01}
type informacoes struct {
	//                     Tag(nome) em json
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	ContaBancaria float64 `json:"ContaBancaria"`
}

func main(){
	// pega as informações de json e transforma em um slice of bytes
	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente não tão secreto","ContaBancaria":8001.01}`)
	
	var jamesbond informacoes
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {fmt.Println(err)}
	
	
	fmt.Println(jamesbond.Nome)
	fmt.Println("\t", jamesbond)

}
//{"Nome":"Manno","Sobrenome":"Stone","Idade":23,"Profissao":"Desenvolvedor?","ContaBancaria":0.06}