package main
import "fmt"

func main(){
	x := struct {
		nome string
		numeros []int
		amigos map[string]string
	}{
		nome: "alfredo",
		numeros: []int{123454,54324},
		amigos: map[string]string{
			"alfredo": "ve ai",
			"henrique": "orbe caolho",}}
	fmt.Println(x)
}
