package components

func NewText(text string, opt ...TextOptions) *TextComponent {

	if len(opt) > 0 {
		return &TextComponent{
			text:        text,
			textOptions: opt[0],
		}
	}
	return &TextComponent{
		text: text,
	}
}

type TextComponent struct {
	text        string
	textOptions TextOptions
}

func (t *TextComponent) RenderSSR() string {

	if t.textOptions != (TextOptions{}) {
		output := "<span style='color:" + t.textOptions.Color + "'>" + t.text + "</span>"
		return output
	}
	

	output := "<span>" + t.text + "</span>"
	return output
}

func (t *TextComponent) String() string {
	return t.text
}
