package main
import ("fmt"
		"sync")

var wg sync.WaitGroup

func main(){
wg.Add(3)
// rotinas
go go1()
go go2()
go go3()
// espera as rotinas terminarem antes de finalizar o programa
wg.Wait()
}

func go1(){
	fmt.Println("Print go1")
	wg.Done()
}
func go2(){
	fmt.Println("Print go2")
	wg.Done()
}
func go3(){
	fmt.Println("Print go3")
	wg.Done()
}
/*
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novasgoroutines(100)
	wg.Wait()
}

func novasgoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número:", i+1)
			wg.Done()
		}(x)
	}
}
*/