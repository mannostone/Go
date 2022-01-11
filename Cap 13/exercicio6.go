package main
import "fmt"
func main(){
	x := 189
	func (x int) {
		fmt.Println(x, "x 545441 =")
		fmt.Println(x*545441)
	}(x)
}