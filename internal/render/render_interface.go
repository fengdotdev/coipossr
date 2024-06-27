package render


// components that can be rendered
type Renderable interface {
	RenderSSR() RenderSSROBJ
}


// the root component must implement this interface
type Cratable interface {
	self() crate
}

type crate struct {
}