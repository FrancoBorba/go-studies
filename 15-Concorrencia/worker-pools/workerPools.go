package main

import "fmt"

/*
Orders come in non-stop: hamburger, pizza, pasta, dessert.
If the manager decided to hire a new cook for each order, the restaurant would collapse: there would be no space,
there would be a shortage of utensils, people would get confused, and the cost would explode.
That's exactly the problem with creating a goroutine for each task.

To avoid this, the restaurant works with a fixed team of cooks.
These cooks stay at their stations and there is an ordering counter. Every new order is placed at that counter.
 As soon as a cook finishes a dish, he takes the next available order and gets to work. When there are no more orders,
  he simply waits. When the restaurant closes, the counter is closed and the cooks close.

This model is the Worker Pool.

In Go, the cooks are the goroutines workers.
The order counter is a channel.
Ready-made dishes can be another channel of results (or not, depending on the problem).

*/
func main() {
	tasks := make(chan int , 45)
	results := make(chan int , 45)

	go worker(tasks , results)
	go worker(tasks , results)
	go worker(tasks , results)
	go worker(tasks , results)

	for i := 0 ; i < 45 ; i++{
		tasks <- i
	} 

	close(tasks)

	for i := 0 ; i < 45 ; i++{
		result := <-results

		fmt.Println(result)
	} 

}

// tasks only send data and result only receive data
func worker(tasks <-chan int , results chan <- int){

	for numero := range tasks {
		results <- fibonacci(numero)
	}
}

func fibonacci(position int ) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position -1)
}