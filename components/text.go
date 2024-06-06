package components

import (
	"github.com/fengdotdev/coipossr/helpers"
	"github.com/fengdotdev/coipossr/opts"
)

func NewText(text string, opt ...opts.TextOptions) *TextComponent {


	if len(opt) > 0 {
		return &TextComponent{
			text:        text,
			textOptions: opt[0],
			id: opt[0].Id,
			className: opt[0].ClassName,
		}
	}
	id := helpers.GenerateId()
	return &TextComponent{
		text: text,
		id: id,
	}
}

type TextComponent struct {
	text        string
	id 		string
	className   string
	tag		string
	textOptions opts.TextOptions
}

func (t *TextComponent) RenderSSR() string {

	if t.textOptions != (opts.TextOptions{}) {
		output := "<span style='color:" + t.textOptions.Color + "'>" + t.text + "</span>"
		return output
	}

	output := "<span>" + t.text + "</span>"
	return output
}

func (t *TextComponent) String() string {
	return t.text
}
