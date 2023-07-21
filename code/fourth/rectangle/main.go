package main

import (
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1024, 1024)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawRectangle(200, 100, 400, 200)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}
