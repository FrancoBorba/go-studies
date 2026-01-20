package main
/*
Imagine now a factory that continuously produces parts.
This factory doesn't produce everything at once and then deliver it; it manufactures it little by little, as someone asks. 
There is a conveyor belt: each time a piece is ready, it is placed on the conveyor belt, 
and whoever is on the other side can pick it up whenever they want. If no one is picking it up, the factory slows down or waits.

That model is the Generator.

In Go, a generator is a goroutine that produces values and sends them over a channel, typically using yield in other languages,
 but in Go this is done with channels + goroutine. 
 The consumer reads these values as he needs to, without knowing (or caring) how they are produced.

The key point of the generator is that it separates production from consumption.
 The consumer does not control the producer's internal loop, they just receive ready-made values. 
 This makes the code simpler, more expressive, and more secure.
*/
import (
	"fmt"
	"time"
)

func main(){
	ch := print("Hello")

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

