package main

import "fmt"

func main(){

	// This channel has a capacity equal to 2 . It will only block when it reaches maximum capacity
	chanel := make(chan string , 2)


	chanel <- "Hello"

	message := <- chanel

	fmt.Println(message)
}