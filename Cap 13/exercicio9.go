package main
import "fmt"
func main() {
	s := quantostem(contanumeros, []int{456,456,456,456,54,6,74868,6587,567,567,5684,846,3}...)
	fmt.Println(s)
}

func contanumeros(x ...int) int {
	return len(x)
}

func quantostem(f func(x ...int) int, i... int) int {
	total := contanumeros(i...)
	return total
}

/*
func main() {
	// chama a função somentePares e passa os valores
	t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	// vai armazenar apenas os inteiros
	var slice []int
	for _, v := range y { // range que vai percorrer
		if v%2 == 0 {
			//slice
			slice = append(slice, v)
		}
	}
	//"callback"
	// total recebe a soma dos valores pares
	// "f(slice...) é a função soma"
	total := f(slice...)
	// valor exibido em main
	return total
}

*/
