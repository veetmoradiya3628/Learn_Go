package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerTime int) int {
    if layerTime == 0 {
        layerTime = 2
    }
    return len(layers) * layerTime;
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 1
		} else if layer == "sauce" {
			sauce += 1
		}
	}
	return noodles * 50, sauce * 0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaleFactor := float64(portions) / 2.0
    scaled := make([]float64, len(quantities))
    for i, q := range quantities {
        scaled[i] = q * scaleFactor
    }
    return scaled
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
