package colors

import "math/rand"

// GetColor returns a color based on the name provided
// if the name is not found, it will return the Black color
func GetColor(name string) Color {
	if _, ok := colorMap[name]; !ok {
		return Black()
	}
	return colorMap[name]
}

// RandomColor returns a random color from the colorMap

func RandomColor() Color {
	keys := colorNames
	randomNumber := randomNumberBetween(0, len(keys))
	randomColor := keys[randomNumber]
	return colorMap[randomColor]
}


func randomNumberBetween(min int, max int) int {
	return min + rand.Intn(max-min)
}

