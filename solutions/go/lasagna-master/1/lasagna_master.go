package lasagnamaster

import "slices"

func PreparationTime(layers []string, avgPrepTime int) int { 
    if (avgPrepTime == 0){
        return len(layers)*2
    } 
    return len(layers) * avgPrepTime
}

func Quantities(layers []string) (noodlesSum int, sauceSum float64) { 
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodlesSum += 50
        } else if layers[i] == "sauce" {
            sauceSum += 0.2
        }
    }
    return
}

func AddSecretIngredient(friendsList []string, myList []string) { 
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

func ScaleRecipe(quantities []float64, portions int) []float64{
    scaled := slices.Clone(quantities) 
    for i := 0; i < len(quantities); i++ {
        scaled[i] *= float64(portions) / 2.0
    }
    return scaled
}

