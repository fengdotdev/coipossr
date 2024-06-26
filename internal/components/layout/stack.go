package layout

import (
	"github.com/fengdotdev/coipossr/components"
	"github.com/fengdotdev/coipossr/render"
)

func Stack () *StackComponent {
	return &StackComponent{}
}


type StackComponent struct {
components []components.Compoment
}

func (s *StackComponent) RenderSSR() render.RenderSSROBJ {
	output := render.RenderSSROBJ{
		JS:   s.renderJS(),
		CSS:  s.renderCSS(),
		HTML: s.renderHTML(),
	}

	return output
}


func (s *StackComponent) AddComponent(com components.Compoment) {
	s.components = append(s.components, com)
}


func (s *StackComponent) renderJS() string {
	return ""
}
func (s *StackComponent) renderCSS() string {
	css:= ""
	return css
}

func (s *StackComponent) renderHTML() string {
	output := ""
	for _, c := range s.components {
		output += c.RenderSSR().HTML
	}
	return output
}

