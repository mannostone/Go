package main
import "fmt"

func main(){
	x := 120
	// y igual a x deslocado em 1
	y := x << 1

	fmt.Printf("%b \n%b\n", x, y)
}