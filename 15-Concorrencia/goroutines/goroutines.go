package main

import (
	"fmt"
	"time"
)

func main(){
	go print("Hello World!")
	go print("Olá Mundo!")

}


func print(value string) {

	for {
		fmt.Println(value)
		time.Sleep(time.Second)
	}
}
