package main

import (
	"fmt"
	"time"
)

func main()  {
	
	i := 0

	for i < 10 {
		i++
		fmt.Printf("The counter value is %d \n" , i)
		time.Sleep(time.Second)
		
		
	}

	for j := 10 ; j < 15; j++ {
		fmt.Printf("The counter of j value is %d \n" , j)
		time.Sleep(time.Second)
	}

	// ------------------- For with the hange clause ------------------

	names := [3]string { "João " , " Pedro" , " Lucas"}

	/*
	The first value will always be the index and the second the value
	*/
	for index , name := range names {
		
		fmt.Printf("The index is %d and the name is %s" , index , name)
	}

	// Ignore the indice
	for _, name := range names {
		
		fmt.Printf("The name  is %s \n", name)
	}


	for index , word := range "FRANCO" {

		fmt.Printf("The char is %s in the index %d \n" , string(word) , index)

	}
}
