package main

import (
	"fmt"
	"tests-introduction/address"
)

func main(){

	typeAdress := address.TypeAdress("Street number 5")

	typeAdress1 := address.TypeAdress("I don't now what is my adress")

	fmt.Println(typeAdress)
	fmt.Println(typeAdress1)
}	