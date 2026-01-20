package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan string) // Defines that this channel can only travel strings
	
	go print("Message from chanel" , chanel)

	for{
	message , open := <- chanel
		if !open {
			break
		}
	fmt.Println(message)
	}

	chanel2 := make(chan string) // Defines that this channel can only travel strings

	go print("Message from chanel 2 " , chanel2)

	// Other way to iterate in a chanel
	for message := range chanel2 {
		fmt.Println(message)

	}

	fmt.Println("End application")
}

func print(value string , chanel chan string) {

	for i:= 0 ; i < 5 ; i++ {
		chanel <- value
		time.Sleep(time.Second)
	}
	close(chanel) // Close the chanel 
}
