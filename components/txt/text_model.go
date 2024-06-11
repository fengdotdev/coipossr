package txt


const DefaultClass = "textComponent"

type TextComponent struct {
	id string
	className string
	text string
	haveOptions bool
	Options     TextOptions
}

func (t *TextComponent) String() string {
	return t.text
}

func (t *TextComponent) Id() string {
	return t.id
}

func (t *TextComponent) ClassName() string {
	return t.className
}

func (t *TextComponent) DefaultClass() string {
	return DefaultClass
}

func (t *TextComponent) FullClass() string {
	return t.DefaultClass() + " " + t.ClassName()
}