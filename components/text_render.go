package components

import (
	"github.com/fengdotdev/coipossr/opts"
	"github.com/fengdotdev/coipossr/types"
)

func (t *TextComponent) RenderSSR() string {
	tag := types.TagSpan()
	tag.SetParam("id", t.id)
	tag.SetParam("class", t.className)
	if t.textOptions != (opts.TextOptions{}) {
		output := "<span style='color:" + t.textOptions.Color + "' >" + t.text + "</span>"
		return output
	}

	output := tag.Start() + t.text + tag.End()
	return output
}