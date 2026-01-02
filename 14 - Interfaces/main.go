package main

import (
	"fmt"
	"math"
)


type form interface{
	area() float64
}

func calculateArea (f form) {
	fmt.Printf("The area is %0.2f \n" , f.area())
}

type rectangle struct  {
	height float64
	large float64
}

func (r rectangle) area() float64{
	return  r.height * r.large
}



type circle struct{
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}
 
func main()  {
	rectangle := rectangle{10,20}
	circle := circle{5}

	calculateArea(rectangle)
	calculateArea(circle)
}