package com

import (
	"github.com/fengdotdev/coipossr/internal/helpers"
	"github.com/fengdotdev/coipossr/internal/types"
)

// constants

const textComponent = "textComponent"
const idPrefix = "text"


//Model

type TextComponent struct {
	id          string
	text        string
	presets     TextPresets
	options     TextOptions
}


//settings
	// public
	type TextOptions struct {
	Id        string
	Groups    []string
	Color     string
	FontStyle string
	FontSize  string
	}
	// private
	type TextPresets struct {
		Groups	[]string
		Color     string
		FontStyle string
		FontSize  string
	}

// constructors

func Text(text string,opts...TextOptions) *TextComponent {
	id := helpers.GenerateIdWithPrefix(idPrefix)
	presets:= TextPresets{}

	options := TextOptions{}
	if len(opts) > 0 {
		options = opts[0]
		id = options.Id
	}

	return &TextComponent{
		id:     id,
		text: text,
		presets: presets,
		options: options,
	}
}



//methods

func (t *TextComponent) String() string {
	return t.text
}

func (t *TextComponent) Id() string {
	return t.id
}

func (t *TextComponent) ClassName() string {
	return 
}

func (t *TextComponent) DefaultClass() string {
	return textComponent
}

func (t *TextComponent) FullClass() string {
	return t.DefaultClass() + " " + t.ClassName()
}

func (t *TextComponent) RenderSSR() types.RenderSSROBJ {

	output := types.RenderSSROBJ{
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

