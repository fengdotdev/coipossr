package com

//Contructors
func Container(child Compoment, opts... ContainerOptions)*ContainerComponent {
	return &ContainerComponent{child: child}
}


//Model
type ContainerComponent struct {
	id      string
	child   Compoment
	presets ContainerPresets
}


//settings
	// public
	type ContainerOptions struct {
		Id        string
		Groups    []string
	}

	// private
	type ContainerPresets struct {
		Groups	[]string
	}




