package main
import "fmt"
// ?
var y interface{}

func main(){
	// Switch com tipos de dados
	y = true
	switch y.(type) {
	case int:
		fmt.Println("inteiro")
	case bool:
		fmt.Println("Bool")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("String")
 
	}
}