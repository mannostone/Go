package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {

	contador := 0
	totaldegoroutines := 10
	// waitgroup
	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func(){
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(contador)
}
