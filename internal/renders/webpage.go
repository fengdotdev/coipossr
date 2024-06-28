package renders

/*

func NewWebPage(r *http.Request, w http.ResponseWriter, c Renderable,opts ...opts.WebPageOptions) *WebPage {
	if len(opts) > 0 {
		return &WebPage{r: r, w: w, component: c,opts:opts[0]}
	}
	return &WebPage{r: r, w: w, component: c}
}

type WebPage struct {
	r         *http.Request
	w         http.ResponseWriter
	component Renderable
	opts 	opts.WebPageOptions
}

func (p *WebPage) Render() error {

	renderObj := p.component.RenderSSR()



	p.w.Header().Set("Content-Type", "text/html")
	p.w.Write([]byte("<!DOCTYPE html>"))
	p.w.Write([]byte("<html>"))
	p.w.Write([]byte("<head>"))
	p.w.Write([]byte("<title>"+p.opts.Title+"</title>"))
	p.w.Write([]byte("<link rel='icon' href='"+p.opts.FaviconUrl+"' type='image/x-icon'>"))
	p.w.Write([]byte(renderObj.Style()))
	p.w.Write([]byte("</head>"))
	p.w.Write([]byte("<body style='background-color:"+p.opts.BackgroundColor+"'>"))
	p.w.Write([]byte(renderObj.Body()))
	p.w.Write([]byte(renderObj.Script()))
	p.w.Write([]byte("</body>"))
	p.w.Write([]byte("</html>"))
	return nil

}

*/