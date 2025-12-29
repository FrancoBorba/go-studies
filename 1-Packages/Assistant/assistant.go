package assistant

import "fmt"

/*
In go functions with the first capital letter no are public. If it is a lowercase, they are private	
*/
/*
In go there is a good practice of commenting on functions that are sporty what it does, 
in this case it would be: Write: Writes a message on the screen
*/
func Writing() {
	fmt.Println("Writing from assitant package")
	writing2()
}