package main
import "fmt"
//tipo indefinido até ser usado em alguma interação
const x = 10
var y = 10

var c int
var d float64

func main(){
	c = x
	fmt.Println(c)
	// x se torna float ja que ele interage com d
	d = x
	fmt.Println(c)
	// Erro pois o tipo de y ja foi definido (linha 5)
	/* d = y
	fmt.Println(y) */

}

/*
// forma alternativa de declarar constantes

const (
	a = 10
	b = 20
	c = 30
)
*/