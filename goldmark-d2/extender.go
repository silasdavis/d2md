package d2

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
)

type Extender struct {
	Layout     d2graph.LayoutGraph
	ThemeID    *int64
	RenderOpts d2svg.RenderOpts
}

func (e *Extender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&Transformer{}, 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&HTMLRenderer{
			Layout:     e.Layout,
			RenderOpts: e.RenderOpts,
		}, 0),
	))
}
