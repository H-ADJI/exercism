package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgpreptime int) (preptime int) {
	if avgpreptime == 0 {
		avgpreptime = 2
	}
	return avgpreptime * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	new_quant := make([]float64, len(quantities))
	for i, v := range quantities {
		new_quant[i] = v * float64(portions) / 2
	}
	return new_quant
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
