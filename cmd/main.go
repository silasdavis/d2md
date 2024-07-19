package main

import (
	d2 "github.com/silasdavis/d2md/goldmark-d2"
	"github.com/yuin/goldmark"
	"os"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

func ptr[T any](v T) *T {
	return &v
}

func main() {
	if len(os.Args) != 2 {
		panic("expected exactly one argument - the markdown filename to process")
	}
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	md := goldmark.New(
		goldmark.WithExtensions(&d2.Extender{
			// Defaults when omitted
			Layout:  d2elklayout.DefaultLayout,
			ThemeID: &d2themescatalog.CoolClassics.ID,
			RenderOpts: d2svg.RenderOpts{
				Scale: ptr(0.5),
			},
		}),
	)
	err = md.Convert(bs, os.Stdout)
	if err != nil {
		panic(err)
	}
}
