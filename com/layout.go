package com

import (
	"github.com/fengdotdev/coipossr/internal/helpers"
)

// Constructors
func Column(childs []Compoment, opts ...LayoutOptions) *LayoutCompoment {
	id := helpers.GenerateId()
	presets:= LayoutPresets{
		LayoutType: ColumnClass,
	}

	options := LayoutOptions{}
	if len(opts) > 0 {
		options = opts[0]
	}

	return &LayoutCompoment{
		id:     id,
		childs: childs,
		presets: presets,
	}
}

// basic layouts


//constants

const LayoutClass = "layout"
const ColumnClass = LayoutClass + "-column"
const RowClass = LayoutClass + "-row"
const StackClass = LayoutClass + "-stack"


//Model
type LayoutCompoment struct {
	id          string
	childs      []Compoment
	presets     LayoutPresets
}

//settings
	// public
	type LayoutOptions struct {
		Id    string
		Groups []string
	}
	// private
	type LayoutPresets struct {
		LayoutType     string
	}




// Methods
