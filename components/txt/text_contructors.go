package txt

import (
	"github.com/fengdotdev/coipossr/helpers"
)

func NewTextWithClass(text string, className string, opts ...TextOptions) *TextComponent {
	id := helpers.GenerateId()
	if len(opts) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			Options: opts[0],
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

func NewTextWithIdAndClass(text string, id string, className string, opts ...TextOptions) *TextComponent {
	if len(opts) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			Options: opts[0],
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

func NewTextWithId(text string, id string, opts ...TextOptions) *TextComponent {
	if len(opts) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			Options: opts[0],
			id:          id,
		}
	}
	return &TextComponent{
		text: text,
		id:   id,
	}
}


func NewText(text string, opts ...TextOptions) *TextComponent {

	id := helpers.GenerateId()
	if len(opts) > 0 {
		return &TextComponent{
			text:        text,
			haveOptions: true,
			Options: opts[0],
			id:          id,
		}
	}
	return &TextComponent{
		text: text,
		id:   id,
	}
}


func Text (text string,opts...TextOptions) *TextComponent {
	return NewText(text,opts...)
}