package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) 

	go func () {
		print("Hello World!")
		waitGroup.Done() // -1

	}()

	go func () {
		 print("Olá Mundo!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0

}


func print(value string) {

	for i:= 0 ; i < 5 ; i++ {
		fmt.Println(value)
		time.Sleep(time.Second)
	}
}
