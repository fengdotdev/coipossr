package types

type Tag struct {
	value  string
	params map[string]string
}

func NewTag(value string) *Tag {
	return &Tag{
		value:  value,
		params: make(map[string]string),
	}
}

func (t *Tag) GetTag() string {
	return t.value
}

func (t *Tag) SetParam(key string, value string) {
	t.params[key] = value
}

func (t *Tag) Start() string {

	params := ""
	for key, value := range t.params {
		params += " " + key + "='" + value + "'"
	}
	return "<" + t.value + params + ">"
}

func (t *Tag) End() string {
	return "</" + t.value + ">"
}

func TagSpan() *Tag {
	return NewTag("span")
}
func TagDiv() *Tag {
	return NewTag("div")
}

func TagH1() *Tag {
	return NewTag("h1")
}

func TagH2() *Tag {
	return NewTag("h2")
}

func TagH3() *Tag {
	return NewTag("h3")
}

func TagH4() *Tag {
	return NewTag("h4")
}

func TagH5() *Tag {
	return NewTag("h5")
}

func TagH6() *Tag {
	return NewTag("h6")
}

func TagP() *Tag {
	return NewTag("p")
}

func TagA() *Tag {
	return NewTag("a")
}

func TagImg() *Tag {
	return NewTag("img")
}

func TagIframe() *Tag {
	return NewTag("iframe")
}

func TagInput() *Tag {
	return NewTag("input")
}

func TagButton() *Tag {
	return NewTag("button")
}
