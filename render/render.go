package render

import "net/http"

func NewPage(w http.ResponseWriter, r *http.Request, component RenderInterface) *Page {
	return &Page{
		r: r,
		w: w,
		component: component,
	}
}

type Page struct {
	r 		*http.Request
	w         http.ResponseWriter
	component RenderInterface
}

func (p *Page)Render() error {
	p.w.Write([]byte(p.component.RenderSSR()))
	return nil
}