package com

import (
	"github.com/fengdotdev/coipossr/internal/helpers"
	"github.com/fengdotdev/coipossr/internal/render"
	"github.com/fengdotdev/coipossr/internal/types"
)

// constants

const textComponent = "textComponent"

// constructors

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

func Text(text string,opts...TextOptions) *TextComponent {
	return NewText(text,opts...)
}


//models

type TextComponent struct {
	id          string
	className   string
	text        string
	haveOptions bool
	Options     TextOptions
}

type TextOptions struct {
	Color     string
	FontStyle string
	FontSize  string
}


//methods

func (t *TextComponent) String() string {
	return t.text
}

func (t *TextComponent) Id() string {
	return t.id
}

func (t *TextComponent) ClassName() string {
	return t.className
}

func (t *TextComponent) DefaultClass() string {
	return textComponent
}

func (t *TextComponent) FullClass() string {
	return t.DefaultClass() + " " + t.ClassName()
}

func (t *TextComponent) RenderSSR() render.RenderSSROBJ {

	output := render.RenderSSROBJ{
		JS:   t.renderJS(),
		CSS:  t.renderCSS(),
		HTML: t.renderHTML(),
	}

	return output
}

func (t *TextComponent) renderJS() string {
	return ""
}

func (t *TextComponent) renderCSS() string {
	css := ""
	return css
}

func (t *TextComponent) renderHTML() string {

	tag := types.TagSpan()
	tag.SetParam("id", t.id)
	tag.SetParam("class", t.FullClass())
	output := tag.Start() + t.text + tag.End()
	return output
}

