package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"
	//                                     valor digitado
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))
	// Quando ok retorna nil
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	// Quando errada, retorna o erro
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}
