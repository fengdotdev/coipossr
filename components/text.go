package components

func NewText(text string) *TextComponent {
	return &TextComponent{
		text: text,
	}
}

type TextComponent struct {
	text string
}

func (t *TextComponent) Render() string {
	return t.text
}

func (t *TextComponent) String() string {
	return t.text
}