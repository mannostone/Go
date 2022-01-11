package main

import "fmt"

func main() {
	//recebe o retorno da função
	x := retornaumafuncao()
	// o retorno da função é aplicado em "y"
	y := x(3)
	fmt.Println(y)
	//forma alternativa de executar a mesma função
	fmt.Println(retornaumafuncao()(3))

}
//   método             anônimo   retorno
func retornaumafuncao() func(int) int {
// retorna uma função anônima
	return func(i int) int {
		return i * 10
	}
}
