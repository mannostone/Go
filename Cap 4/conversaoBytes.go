package main

import "fmt"

func main() {
	a := "e"
	b := "é"
	c := "◙"
	fmt.Printf("%v \n%v\n%v\n", a, b, c)
	// Slice of Bytes
	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v \n%v\n%v\n", d, e, f)
}
