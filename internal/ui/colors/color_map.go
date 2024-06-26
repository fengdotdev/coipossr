package colors

import (
	"fmt"
	"sync"
)


var colorNames []string = func () []string {
	once := sync.Once{}
	var names []string
	once.Do(func() {
		fmt.Println("Creating ColorNames")
		names = ColorNamesSlice()
	})
	return names
}()


func ColorNamesSlice() []string {
	var names []string
	for name := range colorMap {
		names = append(names, name)
	}
	return names
}
func ColorSlices() []Color {
	var colors []Color
	for _, color := range colorMap {
		colors = append(colors, color)
	}
	return colors
}

func ColorMap() map[string]Color {
	return colorMap
}

var colorMap map[string]Color = map[string]Color{

	// White Colors
	"White":     White(),
	"Snow":      Snow(),
	"Honeydew":  Honeydew(),
	"MintCream": MintCream(),
	"Azure":     Azure(),
	"AliceBlue": AliceBlue(),
	"GhostWhite": GhostWhite(),
	"WhiteSmoke": WhiteSmoke(),
	"Seashell": Seashell(),
	"Beige": Beige(),
	"OldLace": OldLace(),
	"FloralWhite": FloralWhite(),
	"Ivory": Ivory(),
	"AntiqueWhite": AntiqueWhite(),
	"Linen": Linen(),
	"LavenderBlush": LavenderBlush(),
	"MistyRose": MistyRose(),

	// Gray Colors
	"Gainsboro": Gainsboro(),
	"LightGray": LightGray(),
	"Silver":    Silver(),
	"DarkGray":  DarkGray(),
	"Gray":      Gray(),
	"DimGray":   DimGray(),
	"LightSlateGray": LightSlateGray(),
	"SlateGray":      SlateGray(),
	"DarkSlateGray":  DarkSlateGray(),
	"Black":     Black(),

}
