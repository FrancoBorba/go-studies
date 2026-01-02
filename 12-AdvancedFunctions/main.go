package main

import "fmt"

/*
In this file I will deal with "Advanced" functions such as functions with named returns, anonymous functions,
recursive functions, functions with pointers, etc. I will choose to keep it in the same file and explicitly document
which one I am using, what its advantages and disadvantages are, etc
*/




func main(){

	sum , subtraction := math(15 , 5)
	fmt.Println("Function with named return")
	fmt.Printf("The result of sum is : %d \n" , sum)
	fmt.Printf("The result of subtraction is : %d \n" , subtraction)
	fmt.Println("---------------------------------------------------------------")


	
	fmt.Println("Variable functions")
	result := infinitySum(1,2,3,4,5)
	fmt.Println("The value of sum in all numbers is " ,result)
	fmt.Println("---------------------------------------------------------------")


	/*
	Anonymous functions in Go are unnamed functions, created at runtime, used when a behavior is local and context-specific. 
	They can be assigned to variables, passed as an argument, executed immediately, or used in goroutines and defer. 
	An important feature is that they can form closures, capturing variables from the outer scope. 
	They are ideal for short, punctual logics, but should be avoided when the function is long, complex,
	 or needs to be reused, as they hinder the readability and testability of the code.
	*/
	fmt.Println("Anonymous functions")

	func (){
		fmt.Println("Hello")
	}()

	var dobro = func(a int) int {
    return a * 2
	}(3) // The parameter

	fmt.Println(dobro)

	fmt.Println("---------------------------------------------------------------")

	fibonnaciPosition := fibonnaci(40)
	fmt.Println(fibonnaciPosition)

	fmt.Println("---------------------------------------------------------------")

	exemploDefer()

	fmt.Println("---------------------------------------------------------------")

	division(3,0)

	fmt.Println("---------------------------------------------------------------")

	fmt.Println("Closure Functions")


	inc := counter()

    fmt.Println(inc()) // 1
    fmt.Println(inc()) // 2
    fmt.Println(inc()) // 3

	fmt.Println("---------------------------------------------------------------")

	numberToInvert := 10
	numberToInvert = invertNumber(numberToInvert)
	fmt.Println(numberToInvert)

	numberToInvert2 := -30
	invertNumberPointer(&numberToInvert2)

	fmt.Println(numberToInvert2)

}

/*
Function with named return 

In this function, variables are already created and initialized in the scope of the function, in this case sum and subtraction. 
This way we can omit it in the return since go understands that it must look for these variables in the scope of the function and return
*/
func math(n1 , n2 int)(sum int , subtraction int){

	sum = n1 + n2
	subtraction = n1 - n2
	return
}

/*
Variable Functions

These are functions that can receive N parameters, without the need to specify these.
In this case we will create a function that adds N numbers
*/

func infinitySum(numbers ... int) (result int) {
	
	for _, value := range numbers {

		result = result + value
	}

	return

}

func fibonnaci(posicao uint ) uint {

	if posicao <= 1{
		return  posicao
	}

	return  fibonnaci(posicao -2) + fibonnaci(posicao - 1)
}


/*
The defer in Go advances the execution of a function to the moment the current function terminates, regardless of where the return occurs. 
Deferred functions are executed in LIFO order (last in, first out). It is widely used to free up resources such as closing files, 
connections, or mutexes, ensuring that cleanup happens even in the event of an error. 
Deferred function arguments are evaluated at the time of defer, not at the end of the function, which avoids unexpected side effects
*/
func exemploDefer() {
    fmt.Println("Início da função")

    // grant schedule this call to be executed
    // automatically when the function returns
    defer fmt.Println("Executa por último")

    fmt.Println("Meio da função")
}


func division(a , b int) int{

	defer func (){

		if r := recover(); r != nil {
			fmt.Println("Erros recover")
		}
	}()

	if b == 0 {
		panic("Division by 0")
	}

	result := a/b

	fmt.Println("Result : " , result)
	return result
}


/*
Functions closure


A closure is a function that captures and maintains access to variables from 
the scope where it was created, even after that scope has already ended.
*/

func counter() func() int {

	cont := 0

	return  func() int {
		cont++
		return  cont;
	}
}


/*
Function with pointers

*/

// With out pointers
func invertNumber(number int) int  {
	return  number * -1
}

// Using pointers
func invertNumberPointer(number *int) {
	*number = *number * -1
}
