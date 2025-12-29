package main

import "fmt"

type person struct{
	name string
	lastname string
	age int
	
}

type student struct{
	person // "Inheritance"
	course string
	colegge string
	enrollment float64
}

func main(){
	 p1 := person{"Joao" , "Pedro" , 56}
	fmt.Println(p1)

	student1 := student{p1 , "Computação" , "UESB" , 202510445}

	fmt.Println(student1)
	fmt.Println(student1.name) // We can take the name of student
}
