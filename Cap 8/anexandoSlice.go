package main

import "fmt"

func main() {
	umaslice := []int{1,2,3,4} 
	outraslice := []int{5,6,7,8}

	umaslice = append(umaslice,9,10,11,12)
	// "outraslice..." passa os valores do vetor.
	// o append espera receber valores do tipo da slice e não uma outra slice de inteiros
	umaslice = append(umaslice[:], outraslice...)

	fmt.Println(umaslice)
}
