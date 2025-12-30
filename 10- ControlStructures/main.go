package main

import "fmt"

func main()  {
	numero1 := 6

	/*
	The if and else structures in go are practically the same as in other languages like java, 
	so I won't spend my time creating if and else examples.
	We can highlight as a difference the non-obligation to put "()" in the condition
	*/
	if(numero1 >= 5 && numero1 <= 10){
		fmt.Println("Number is beetween 5 and 10")
	} else{
		fmt.Println("Number isn´t beetween 5 and 10")
	}

	
	// As a differential we can highlight the if init , which would be to create a variable in the condition block
	// The variable is local in the block
	if createNumber := 100; createNumber > 50 {
		fmt.Println("The number is bigger than 50")
	}

	if createNumber := numero1; createNumber > 50 {
		fmt.Println("The number is bigger than 50")
	}else{
		fmt.Println("The number isn´t bigger than 50")
	}

	// --------------- SWITCH IN GO --------------

	var day uint8 
	day = 4
	fmt.Println(dayOfWeek(day))
	
}

	func dayOfWeek(day uint8) string{
		switch day {
		case 1: 
			return  "Domingo"
		case 2:
			return  "Segunda"
		case 3:
			return  "Terça"
		case 4:
			return  "Quarta"			
		
		case 5:
			return "Quinta"
		case 6: 
			return "Sexta"
		case 7: 
			return  "Sabado"
		default:
			return "Number is invalid"
		}
	}

	
	func evaluateGrade(grade uint8) string {
	switch {
	case grade >= 90:
		{
			return "Excelente"
		}
	case grade >= 75:
		{
			return "Bom"
		}
	case grade >= 60:
		{
			return "Satisfatório"
		}
	case grade >= 50:
		{
			return "Aprovado"
		}
	default:
		{
			return "Reprovado"
		}
	}
}