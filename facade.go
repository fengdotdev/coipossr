package coipossr

import (
	"net/http"

	"github.com/fengdotdev/coipossr/components"
	"github.com/fengdotdev/coipossr/render"
)

//renders
func Page(w http.ResponseWriter, r *http.Request, component render.RenderInterface) *render.WebPage {
	return render.NewWebPage(r, w, component)
}

// Components
func Text(text string) *components.TextComponent {
	return &components.TextComponent {
		Text: text,
	}
}