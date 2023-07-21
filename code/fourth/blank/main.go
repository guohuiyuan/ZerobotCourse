package main

import "github.com/fogleman/gg"

func main() {
	dc := gg.NewContext(100, 100)
	dc.SavePNG("out.png")
}
