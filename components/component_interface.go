package components

import (
	"github.com/fengdotdev/coipossr/render"
	"github.com/fengdotdev/coipossr/self"
)

type Compoment interface {
	render.RenderInterface
	self.SelfInterface
}