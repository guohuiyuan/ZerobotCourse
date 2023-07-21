package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(200, 100)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SavePNG("out.png")
}