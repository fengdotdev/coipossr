package layout

import "github.com/fengdotdev/coipossr/components"

func Column(components []components.Compoment, opts ...LayoutOptions) *ColumnCompoment {
	return &ColumnCompoment{
		className:  DefaultClass,
		components: components,
	}
}