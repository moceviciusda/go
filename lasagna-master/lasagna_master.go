package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgLayerPrepTime int) int {
	if avgLayerPrepTime == 0 {
		avgLayerPrepTime = 2
	}
	return len(layers) * avgLayerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsReciepe []string, myReciepe []string) {
	myReciepe[len(myReciepe)-1] = friendsReciepe[len(friendsReciepe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := []float64{}
	for _, amount := range amounts {
		scaledAmounts = append(scaledAmounts, amount / 2 * float64(portions))
	}

	return scaledAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
