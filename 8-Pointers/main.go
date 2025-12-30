package main

import "fmt"

func main()  {
	
	var variable1 int = 10;
	var variable_withoutPointer int = variable1

	fmt.Println(variable1, variable_withoutPointer)

	variable1++;
	fmt.Println(variable1, variable_withoutPointer)


	var variable2 int = 100
	var varible_with_pointer *int = &variable2 // Memory reference &

	fmt.Println(variable2 , varible_with_pointer)
	fmt.Println(variable2 , *varible_with_pointer) // In this way we see the value and not the memory adress

	variable2 += 10

	fmt.Println(variable2 , *varible_with_pointer) 



}
