package main
import "fmt"

const (
	_ = iota + 2021
	a
	b
	c
	d
)
func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}