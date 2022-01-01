package main

import "fmt"

func main() {
	amigos := map[string]int{
		"Alfredo": 551234543,
		// 12 key phone layout - 
		// "6-666-777-777-2-66-444-66-8-99-66-9-666"
		"Nintendo": 66667777772,
	}

	fmt.Println(amigos)

	amigos["Void"] = 0000000000
	fmt.Println(amigos)
	fmt.Println(amigos["Void"])

}
