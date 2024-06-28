package builders

import (
	"github.com/fengdotdev/coipossr/internal/crates"
	"github.com/fengdotdev/coipossr/internal/routers"
)	



// constructor

func NewAppWasm(pages []crates.WasmPage,router routers.WasmRouter) *AppWasm {
	return &AppWasm{
		pages: make(map[string]crates.WasmPage),
		router: router,
	}
}

type AppWasm struct {
	pages map[string]crates.WasmPage
	router routers.WasmRouter
}


func (app *AppWasm) Run () error {
	return nil
}

func 

func (app *AppWasm) AddPage(page crates.WasmPage) error{
	return nil
}	