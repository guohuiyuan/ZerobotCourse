package main

import (
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1024, 1024)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("code/fourth/regular-bold.ttf", 50); err != nil {
		panic(err)
	}
	dc.DrawString("你好", 0, 50)
	dc.SavePNG("out.png")
}
