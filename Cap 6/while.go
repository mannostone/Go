package main
import "fmt"
// Não existe o while em Go
func main(){
	//while "camuflado"
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// "break" quebra o loop 
	// "continue" quebra a interação atual do loop
	// Exemplo:
	/* 
	for i :=0; i<10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
	*/
}