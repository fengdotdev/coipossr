package colors

func NewColor(name string, red int, green int, blue int, transparency float64, hex string) *Color {
	return &Color{
		Name:         name,
		Red:          red,
		Green:        green,
		Blue:         blue,
		Transparency: transparency,
		Hex:          hex,
	}
}

type Color struct {
	Name         string
	Red          int
	Green        int
	Blue         int
	Transparency float64
	Hex          string
}