package main

import "fmt"

func main()  {
	
	// Inside the bracket is the key type outside the bracket is the values type
	// In this case it would be numbers for the key and names for the values 
	//For example Franco is associated with 1
	user := map[uint]string{
		1: "Franco",
		2: "Adryellen",
		3: "João",
	}
	
	fmt.Println(user)
	fmt.Println(user[1])

	delete(user , 1) // delete the ey and value
	fmt.Println(user)

	




}
