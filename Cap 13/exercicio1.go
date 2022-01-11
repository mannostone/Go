package main

import "fmt"

func main(){
	si := []int {10,10,1,2,3,4,5,6,8,9,23}
	// forma de passar slices como argumentos
	total := soma(si...)

	fmt.Println(total)

	numero, letra := (intstring(20,"arroz"))
	fmt.Println(numero,letra)

}
//   nome(assinatura variádica) retorno {}
func soma(x... int) int {
	soma := 0
	for _, v := range x{
		soma+=v
	}
	return soma
}

func intstring (i int, s string) (int,string){
	return i,s
}