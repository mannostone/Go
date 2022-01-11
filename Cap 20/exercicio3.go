package main
import ("fmt"
		"runtime"
		"sync")

var contador = 0
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
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}