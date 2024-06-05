package render

import "net/http"


type PageInterface interface {
	Render() error
}

func NewWebPage(r *http.Request, w http.ResponseWriter, c RenderInterface) *WebPage {
	return &WebPage{r: r, w: w, component: c}
}


type WebPage struct {
	r 		*http.Request
	w         http.ResponseWriter
	component RenderInterface
}

func (p *WebPage)Render() error {
	p.w.Write([]byte(p.component.RenderSSR()))
	return nil
}