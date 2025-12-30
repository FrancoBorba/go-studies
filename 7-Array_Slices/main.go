package main

import (
	"fmt"
	"reflect"
)

func main()  {
	
	// Create a array of 5 position 
	var array[5] int
	fmt.Println(array)

	array[0] = 3
	array[1] = 6

	fmt.Println(array)

	// Fill in the first 3 positions and leave the others at default value
	array1 := [5]string {"Posision 1" , "Positions 2 ", "Test "}

	fmt.Println(array1)

	// -------------- SLICE --------------------

	slice := []int{1,34,556,33,2,9}
	fmt.Println(slice)
	
	fmt.Println(reflect.TypeOf(slice) )
	fmt.Println(reflect.TypeOf(array) )

	slice = append(slice, 6)

	fmt.Println(slice)

	/*
	
	It points to array 2, it's like a pointer. 
	The first parameter is inclusive, that is, it starts from it. 
	The second parameter is exclusive, that is, it goes to the previous one in this case, it goes from 0 to 1 
	*/
	slice2 := array[0:2]
	fmt.Println(slice2)

	slice2 = append(slice2, 9 , 12)

	fmt.Println(slice2)

	array[0] = 1 // will change in the slice
	fmt.Println(slice2)

	// Internal Arrays

	slice3 := make([]float32 , 10 , 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Lenght
	fmt.Println(cap(slice3)) // capacity



}
