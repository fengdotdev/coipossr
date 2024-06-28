package renders

import "github.com/fengdotdev/coipossr/internal/types"

// the childs components must implement this interface
type Renderable interface {
	RenderSSR() types.RenderSSROBJ
}

//Page
type Cratable interface {
	self() types.Crate
}

