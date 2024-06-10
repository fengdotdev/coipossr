package layout

import "github.com/fengdotdev/coipossr/components"

type Column struct {
components []components.Compoment
}


func (c *Column) RenderSSR() string {
	var html string

	for _, c := range c.components {
		html += c.RenderSSR()
	}
	return html
}