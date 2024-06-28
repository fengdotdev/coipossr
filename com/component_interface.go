package com

import "github.com/fengdotdev/coipossr/internal/renders"

// components must be inmutable and must implement the Renderable Interface and SelfInterface


type Compoment interface {
	renders.Renderable
	self()componentizable
}


type componentizable struct {
}