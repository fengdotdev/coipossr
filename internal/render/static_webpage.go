package render

import (
	"io"

	"github.com/fengdotdev/coipossr/opts"
)

func NewStaticWebPage(c Renderable,opts ...opts.WebPageOptions) *StaticWebPage {
	if len(opts) > 0 {
		return &StaticWebPage{component: c,opts:opts[0]}
	}
	return &StaticWebPage{ component: c}
}

type StaticWebPage struct {
	component Renderable
	opts 	opts.WebPageOptions

}

func (p *StaticWebPage) Render(w io.Writer) error {

	renderObj := p.component.RenderSSR()



	//p.w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<!DOCTYPE html>"))
	w.Write([]byte("<html>"))
	w.Write([]byte("<head>"))
	w.Write([]byte("<title>"+p.opts.Title+"</title>"))
	w.Write([]byte("<link rel='icon' href='"+p.opts.FaviconUrl+"' type='image/x-icon'>"))
	w.Write([]byte(renderObj.Style()))
		w.Write([]byte("</head>"))
	w.Write([]byte("<body style='background-color:"+p.opts.BackgroundColor+"'>"))
	w.Write([]byte(renderObj.Body()))
	w.Write([]byte(renderObj.Script()))
	w.Write([]byte("</body>"))
	w.Write([]byte("</html>"))
	return nil

}