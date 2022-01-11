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
	// mutex
	var mu sync.Mutex 

	for i := 0; i < totaldegoroutines; i++ {
		go func(){
			mu.Lock() // 
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock() // 
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(contador)
}
