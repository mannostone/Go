package main
import "fmt"

func main(){
	ano := 1998
	for {
		if ano - 1998 == 0 {
			fmt.Println("Nasci")
		} else if ano - 1998 == 1{
			fmt.Printf("%v ano de vida\n", ano-1998)
		} else {
		fmt.Printf("%v anos de vida\n", ano-1998)
		}

		if (ano == 2021){
			break
		}
		ano++
	}
}