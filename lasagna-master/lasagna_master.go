package lasagna


const (
    defaultTime = 2 //in minutes
    saucePerLayer = .2 //in liters
    noodlesPerLayer = 50 //in grams
    defaultServingCount = 2
)
func PreparationTime(slice []string, prepTime int) int {
	if prepTime <= 0 {
        totalTime := (len(slice) * defaultTime)
        return totalTime
    } else {
    	realPrepTime := prepTime
        totalTime := (len(slice) * realPrepTime)
        return totalTime
    }
}


func Quantities(slice []string) (noodles int, sauce float64) {
    for _, layers := range slice {
		switch layers {
            case "noodles":
        		noodles++
            case "sauce":
        		sauce++
        }
    }
	
	return noodles * noodlesPerLayer, float64(sauce) * saucePerLayer
}

func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portion int) []float64 {
    var adjustedQuantities []float64
    for _, recipeQty := range quantities {
        adjustedQuantities = append(adjustedQuantities, recipeQty / defaultServingCount * float64(portion))
    }
    return adjustedQuantities
}