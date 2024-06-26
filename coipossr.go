package coipossr

import (
	"net/http"

	"github.com/fengdotdev/coipossr/opts"
	"github.com/fengdotdev/coipossr/render"
)

//renders

func StaticPage(component render.RenderInterface,opts ...opts.WebPageOptions) *render.StaticWebPage {
	return render.NewStaticWebPage(component,opts...)
}

func Page(w http.ResponseWriter, r *http.Request, component render.RenderInterface,opts ...opts.WebPageOptions) *render.WebPage {
	return render.NewWebPage(r, w, component,opts...)
}

// Components
