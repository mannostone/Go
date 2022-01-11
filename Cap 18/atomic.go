package main

import (
	"fmt"
	"sync/atomic"
	"runtime"
)

func main() {

	var contador int64
	contador = 0
	totaldegoroutines := 10
	// waitgroup
	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func(){
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(contador)
}
