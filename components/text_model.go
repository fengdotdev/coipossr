package components

import (
	"github.com/fengdotdev/coipossr/opts"
)



const TextComponentClassName = "textComponent"



type TextComponent struct {
	text        string
	id 		string
	className   string
	haveOptions bool
	textOptions opts.TextOptions
}


func (t *TextComponent) String() string {
	return t.text
}

func (t *TextComponent) Id () string {
	return t.id
}

func (t *TextComponent) ClassName() string {
	return t.className
}
