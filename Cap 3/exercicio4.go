package main

import "fmt"

type HK int
var x HK

func main(){
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Println(x)

}