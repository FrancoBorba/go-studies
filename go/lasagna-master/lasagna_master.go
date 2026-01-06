package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string , preparationTimePerLayer int) int   {
	
	if(preparationTimePerLayer == 0){
		preparationTimePerLayer = 2
	}

	return  len(layers) * preparationTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int ,float64){

	var amountNoodles int
	var amountSauce float64

	for _ , layer := range layers{

		switch layer {
			
		case "noodles":
			amountNoodles = amountNoodles + 50
		case "sauce":

			amountSauce = amountSauce + 0.2
		}
	}
	return  amountNoodles, amountSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

	scaleFactor := float64(portions) / 2
	scaled := make([]float64, len(quantities))

	for i, quantity := range quantities {
		scaled[i] = quantity * scaleFactor
	}

	return scaled
}

