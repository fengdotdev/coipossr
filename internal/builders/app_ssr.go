package builders

import (
	"github.com/fengdotdev/coipossr/internal/crates"
)

type AppSSR struct {
	pages map[string]crates.SSRPage
}

func (app AppSSR) AddPage(page crates.SSRPage) error{
	//TODO: add validation
	app.pages["name"] = page
	return nil
}