package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %t\n", a,a)
	fmt.Printf("%v, %t\n", b,b)
	fmt.Printf("%v, %t\n", c,c)
	fmt.Printf("%v, %t\n", d,d)
}