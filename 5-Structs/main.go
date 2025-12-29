package main

import "fmt"


type user struct {
	name string
	age uint8
	address address
}

type address struct{
	street string
	number int
}
func main(){

	var user_franco user
	user_franco.name = "Franco"
	user_franco.age = 22
	fmt.Println(user_franco)

	addres_pedro := address{"Olivia" , 1005}

	user_pedro := user {"Jonas" , 34 , addres_pedro}
	fmt.Println(user_pedro)

	// In this way we can pass just the arguments wanteds
	user_cassia := user{name: "Cassia"}
	fmt.Println(user_cassia)


}