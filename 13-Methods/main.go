package main

import "fmt"

//A method is a function associated with a type.
// The difference to an ordinary function is that the method has a receiver.
type user struct{
	name string
	age uint8
}

// The method here is associated a user
func (u user) salvar(){
	fmt.Printf("Saving the user %s \n" , u.name)
}

func (u user) checkIfUserIsAdult () bool{

	if(u.age >= 18) {
		return  true
	} 

	return  false
}

// Need be a pointer to 
/*
 Change the original struct
 Avoids copying large structs

 Rule of thumb:
If the method modifies the state, use pointer.
*/
func (u *user) makeBirthday (){
	u.age ++
}

func main(){

	user1 := user{"Franco" , 22}
	var user2 user
	user2.name = "Greco"

	user1.salvar()
	user2.salvar()

	if (user1.checkIfUserIsAdult()){
		fmt.Println("Is an adult")
	} else{
		fmt.Println("Isnt an adult")
	}

	user2.age = 17

	if (user2.checkIfUserIsAdult()){
		fmt.Println("Is an adult")
	} else{
		fmt.Println("Isnt an adult")
	}

	user2.makeBirthday()

		if (user2.checkIfUserIsAdult()){
		fmt.Println("Is an adult")
	} else{
		fmt.Println("Isnt an adult")
	}

}
