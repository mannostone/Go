package main
import "fmt"

const (
	// tipos indefinidos!
	x = iota
	y = iota
	z = iota
)

func main(){
	// exibição dos tipos indefinidos (0,1,2)
	fmt.Println(x, y, z)
}