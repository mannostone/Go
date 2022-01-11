package main

import (
	"fmt"
	"runtime"
)



func main(){
	fmt.Println("Seu OS é:\n", runtime.GOOS)
	fmt.Println("Seu ARCH é:\n", runtime.GOARCH)
}