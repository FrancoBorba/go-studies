package main

import (
	"fmt"
	"time"
)

func main(){
	ch := multiplexar(print("Hello") , print("World"))

	for i := 0 ; i < 10 ; i++ {
		fmt.Println(<-ch)
	}
}

func print(text string) <- chan string{
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Value reciver: %s" , text)
			time.Sleep(time.Second)
		}
	}()

	return  ch
}

func multiplexar ( inputChanel1 <-chan string , inputChael2 <-chan string) <- chan string {

	outputChanel := make(chan string )

		go func() {
		for {
			select {
			case message := <- inputChanel1: 
				outputChanel <- message
			case message := <- inputChael2:
				outputChanel <- message
			}
		}
	}()
	return  outputChanel
}