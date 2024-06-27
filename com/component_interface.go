package com

import (
	"github.com/fengdotdev/coipossr/internal/render"
)

// components must be inmutable and must implement the Renderable Interface and SelfInterface


type Compoment interface {
	render.Renderable
	self()componentizable
}


type componentizable struct {
}