package components

func NewText(text string) *TextComponent {
	return &TextComponent{
		text: text,
	}
}

type TextComponent struct {
	text string
}

func (t *TextComponent) RenderSSR() string {
	output := "<span>" + t.text + "</span>"
	return output
}

func (t *TextComponent) String() string {
	return t.text
}