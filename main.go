package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pandatix/gojs-cvss/internal"
)

func main() {
	js.Global.Set("CVSS40", internal.New40)
	js.Global.Set("Rating", internal.Rating)

	if js.Module != js.Undefined {
		exp := js.Module.Get("exports")
		exp.Set("cvss40", map[string]any{
			"New": internal.New40,
		})
		exp.Set("Rating", internal.Rating)
	}
}
