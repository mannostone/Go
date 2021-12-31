package main
import "fmt"

func main(){
	x := (1 == 1)
	fmt.Println(x)
	x = (1 != 1)
	fmt.Println(x)
	x = 1 <= 2
	fmt.Println(x)
	x = 1 < 2
	fmt.Println(x)
	x = 1 >= 2
	fmt.Println(x)
	x = 1 > 2
	fmt.Println(x)

}