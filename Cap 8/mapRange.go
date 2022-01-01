package main

import "fmt"

func main() {
	
	qualquerCoisa := map[int]string{
		123: "Legal hein",
		98: "éééé...",
		983: "aceitável",
	}

	fmt.Println(qualquerCoisa)
	
	for key, value := range qualquerCoisa {
		fmt.Println(key, " ", value)
	}

	total := 0

	for key, _ := range qualquerCoisa {
		total += key
	}

	fmt.Println(total)

	delete(qualquerCoisa, 123)
	fmt.Println(qualquerCoisa)


}
