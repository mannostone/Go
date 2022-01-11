package main
import ("fmt"
		"runtime"
		"sync/atomic")

var contador =int64 (0)
var wg sync.WaitGroup
const routines = 5

func main(){
	criarGoroutines(routines)
	wg.Wait()
	fmt.Println("Total de goroutines: ", routines)
	fmt.Println("Total do contador: ", contador)
}

func criarGoroutines (i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func () {
			//              endereço / valor acrescentado
			atomic.AddInt64(&contador, 1)
			v := atomic.LoadInt64(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}