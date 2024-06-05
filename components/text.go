package components

func Text(text string) *TextComponent {
	return &TextComponent{
		Text: text,
	}
}

type TextComponent struct {
	Text string
}

func (t *TextComponent) Render() string {
	return t.Text
}

func (t *TextComponent) String() string {
	return t.Text
}