package main

import (
	"fmt"
	"modulo/Assistant"

	"github.com/badoux/checkmail"
)

func main () {

	fmt.Println("Writing from main file")
	assistant.Writing()

	erro := checkmail.ValidateFormat("franco.borba	@gmail")
	fmt.Println(erro)
}

