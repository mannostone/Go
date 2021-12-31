package main
import "fmt"

const x int = 10
const y = 10

func main(){
	fmt.Printf("%v\n%T\n%v\n%T\n", x,x,y,y)
}