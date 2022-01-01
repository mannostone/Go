package main
import "fmt"

var s string

func main(){
	for i := 33; i < 123; i++ {
		// exibindo os valores no formato string(texto Ascii)
		fmt.Printf("%v\n", string(i))
	}
	
}