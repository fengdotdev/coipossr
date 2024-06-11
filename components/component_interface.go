package components

import (
	"github.com/fengdotdev/coipossr/render"
	"github.com/fengdotdev/coipossr/self"
)

// components must be inmutable and must implement the RenderInterface and SelfInterface


type Compoment interface {
	render.RenderInterface
	self.SelfInterface
}


