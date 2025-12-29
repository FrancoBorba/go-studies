package main

import "fmt"

func main()  {
	var name string = "Name 1" // Explicitly declaring
	name2 := "Name 2"  // Type inference

	fmt.Println(name)
	fmt.Println(name2)

	var (
		name3 string = "Name 3"
		name4 string = " Name 4"
	)

	fmt.Println(name3 , name4)


	name5 , name6 := "Name 5" , "Name 6"

	fmt.Println(name5 , name6)
}
