package render

import (
	"io"

	"github.com/fengdotdev/coipossr/types"
)

type RenderSSROBJ struct {
	JS   string
	CSS  string
	HTML string
}

func (r *RenderSSROBJ) String() string {
	return r.HTML
}

func (r *RenderSSROBJ) Merge(obj RenderSSROBJ){
	r.JS += obj.JS
	r.CSS += obj.CSS
	r.HTML += obj.HTML
}
func (r *RenderSSROBJ) MergeCSS(css string){
	r.CSS += css
}

func (r *RenderSSROBJ) MergeJS(js string){
	r.JS += js
}

func (r *RenderSSROBJ) MergeHTML(html string){
	r.HTML += html
}


func (r *RenderSSROBJ) Script() string {
	tag := types.TagScript()
	output := tag.Start() + r.JS + tag.End()
	return output
}

//Inline CSS
func (r *RenderSSROBJ) Style() string {
	tag := types.TagStyle()
	output := tag.Start() + r.CSS + tag.End()
	return output
}

func (r *RenderSSROBJ) ExternalCSS(w io.Writer)error{
	//TODO
	panic("Not implemented")
	return nil
}

func (r *RenderSSROBJ) Body() string {
	return r.HTML
}
