package address

import "strings"

// TypeAdress checks if it has a valid address and returns
func TypeAdress(adress string) string {
	validTypes := []string{"street" , "avenue" , "highway" , "neighborhood"}

	firstWord := strings.Split(strings.ToLower(adress), " ")[0]

	for _ , types := range validTypes{
		if(types == firstWord) {
			return firstWord
		}
	}

	return "Invalid Type"
}