package main

import "fmt"

func main()  {
	result := sum(8, 23)
	fmt.Println(result)

	resultMultiply , resultDivisio := multipliANDdivision(100,10)
	fmt.Println(resultMultiply, resultDivisio)

	// If you juust want one result , you can use a _ to replace a variable

	resultMultiply2 , _ := multipliANDdivision(10,10)
	fmt.Println(resultMultiply2)

	_ , resultDivision2 := multipliANDdivision(1,10)

	fmt.Println(resultDivision2)

}

func sum(number1 int , number2 int) int{
	return  number1 + number2
}
/*
In go we can have more than one return per fuunction. 
It is extremely used mainly in error handling scenes

*/
func multipliANDdivision(n1, n2 int16) (int16, int16){


	multiply := n1 * n2	
	division := n1/n2

	return multiply, division

}