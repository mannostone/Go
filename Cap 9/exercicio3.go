package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(slice[0:3])
	fmt.Println(slice[4:len(slice)])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:(len(slice)-1)])
	
}
/* 
[1 2 3 4 5 6 7 8 9 10]

[1 2 3]
[5 6 7 8 9 10]
[2 3 4 5 6 7]
[3 4 5 6 7 8 9]
*/