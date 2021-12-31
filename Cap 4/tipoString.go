package main
import "fmt"

func main(){
	s := "Aloo"

	// Exibição de caracter por caracter
	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println()
	// Exibição de byte por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}