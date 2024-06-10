package layout

import "github.com/fengdotdev/coipossr/components"

func Stack () *StackComponent {
	return &StackComponent{}
}


type StackComponent struct {
components []components.Compoment
}

func (s *StackComponent) RenderSSR() string {
	var html string
	
	for _, c := range s.components {
		html += c.RenderSSR()
	}	
	return html
}


func (s *StackComponent) AddComponent(c components.Compoment) {
	s.components = append(s.components, c)
}

