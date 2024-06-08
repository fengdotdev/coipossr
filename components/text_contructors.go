package components

import (
	"github.com/fengdotdev/coipossr/helpers"
	"github.com/fengdotdev/coipossr/opts"
)

func NewTextWithClass(text string, className string, opt ...opts.TextOptions) *TextComponent {
	id := helpers.GenerateId()
	if len(opt) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			textOptions: opt[0],
			className:   className,
			id:          id,
		}
	}
	return &TextComponent{
		text:      text,
		className: className,
		id:        id,
	}
}

func NewTextWithIdAndClass(text string, id string, className string, opt ...opts.TextOptions) *TextComponent {
	if len(opt) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			textOptions: opt[0],
			id:          id,
			className:   className,
		}
	}
	return &TextComponent{
		text:      text,
		id:        id,
		className: className,
	}
}

func NewTextWithId(text string, id string, opt ...opts.TextOptions) *TextComponent {
	if len(opt) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			textOptions: opt[0],
			id:          id,
		}
	}
	return &TextComponent{
		text: text,
		id:   id,
	}
}

func NewText(text string, opt ...opts.TextOptions) *TextComponent {

	id := helpers.GenerateId()
	if len(opt) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			textOptions: opt[0],
			id:          id,
		}
	}
	return &TextComponent{
		text: text,
		id:   id,
	}
}