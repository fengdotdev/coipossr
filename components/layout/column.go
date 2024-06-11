package layout

import (
	"github.com/fengdotdev/coipossr/components"
	"github.com/fengdotdev/coipossr/render"
)










type ColumnCompoment struct {
	id string
	className string
	components []components.Compoment
	haveOptions bool
	Options     LayoutOptions
}


func (c *ColumnCompoment) RenderSSR() render.RenderSSROBJ {
	output := render.RenderSSROBJ{
		JS: c.renderJS(),
		CSS: c.renderCSS(),
		HTML: c.renderHTML(),
		}
	return output
}

func (c *ColumnCompoment) AddComponent(com components.Compoment) {
	c.components = append(c.components, com)
}

func (c *ColumnCompoment) renderJS() string {
	return ""
}

func (c *ColumnCompoment) renderCSS() string {
	css := ""
	return css
}

func (c *ColumnCompoment) renderHTML() string {
	output := ""
	for _, c := range c.components {
		output += c.RenderSSR().HTML
	}
	return output
}


func (c *ColumnCompoment) SelfCSS ()string {

	css := "." + DefaultClass + "{"
	css += "display: flex;"
	css += "flex-direction: column;"
	css += "}"
	return css
}


