package com

//constants

const Layo = "layout"
const ColumnClass = DefaultClass + "-column"
const RowClass = DefaultClass + "-row"
const StackClass = DefaultClass + "-stack"


//Models
type LayoutCompoment struct {
	id          string
	defaultClass string
	className   string
	components  []Compoment
	haveOptions bool
	Options     LayoutOptions
}

type LayoutOptions struct {
}

// Constructors




// Methods
func (l *LayoutCompoment) String() string {
	return l.className
}

func (l *LayoutCompoment) Id() string {
	return l.id
}

func (l *LayoutCompoment) ClassName() string {
	return l.className
}

func (l *LayoutCompoment) DefaultClass() string {
	return l.defaultClass
}

func (l *LayoutCompoment) FullClass() string {
	return l.defaultClass + " " + l.ClassName()
}
