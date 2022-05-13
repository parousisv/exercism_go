package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparation_time int) int {
    if (preparation_time > 0){
        return preparation_time * len(layers)
    }
    return 2 * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauces := 0.0
    for i := 0; i < len(layers); i++{
        if layers[i] == "noodles"{
            noodles++
        }else if layers[i] == "sauce"{
			sauces++
        }
    }
	return noodles*50, sauces*0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaled_quantities := make([]float64, len(quantities))
	for i := 0; i < len(scaled_quantities); i++{
        scaled_quantities[i] = (quantities[i]/2.0) * float64(portions)
    }
	return scaled_quantities
}
