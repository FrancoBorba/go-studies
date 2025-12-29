package main

import (
	"errors"
	"fmt"
)

func main()  {

	// Numbers

	// ----------------  ints --------------

	var (
	 number1 int8 = 8 // 8 bits
	 number2 int16 = 16 // 16 bits
	 number3 int32 = 32 // 32 bits
	 number4 int64 = 64 // 64 bits
	 number5 int = 100 // It will use the computer architecture
	)

	fmt.Println(number1 , number2 , number3 , number4 , number5)

	// We can specify the bits as the int
	var number6 uint = 9 // Unsigned numbers

	fmt.Println(number6)


	// ------- Float ----------

	var number7 float32 = 14.6
	var number8 float64 = -4139.5432

	fmt.Println(number7 , number8)


	// --------- Strings -----------

	var str string = "Texto";
	texto2 := "Texto2";

	fmt.Println(str , texto2);

	// As it is in single quotes it would be the char, in go there is no char and it fetches the asc table and assigns the corresponding value
	char := 'b' 

	fmt.Println(char)


	// -------------- Boolean ----------

	// Just true or false
	var IamBeauty bool = true

	fmt.Println(IamBeauty)

	// ------------ Error --------------

	var erro error = errors.New("Intern erro")
	fmt.Println(erro)
}
