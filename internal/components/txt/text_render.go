package txt

import (
	"github.com/fengdotdev/coipossr/render"
	"github.com/fengdotdev/coipossr/types"
)

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
	css:= ""
	return css
}

func (t *TextComponent) renderHTML() string {
	
	tag := types.TagSpan()
	tag.SetParam("id", t.id)
	tag.SetParam("class", t.FullClass())
	output := tag.Start() + t.text + tag.End()
	return output
}

