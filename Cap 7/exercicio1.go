package main
import "fmt"

func main(){
	for i := 1; i <=10000; i++{
		fmt.Printf("%v ", i)
	
		if (i%30 == 0){
			fmt.Println()
		}
	}
}