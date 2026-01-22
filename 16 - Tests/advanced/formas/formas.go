package formas

import (
	"math"
)


type Form interface{
	Area() float64
}


type Rectangle struct  {
	Height float64
	Large float64
}

func (r Rectangle) Area() float64{
	return  r.Height * r.Large
}



type Circle struct{
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius,2)
}
 
