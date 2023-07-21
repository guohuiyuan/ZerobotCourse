package main

import (
	"fmt"
	"image"
	"os"

	"github.com/FloatTech/imgfactory"
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(4000, 2000)
	back, err := gg.LoadImage("fourth/原神.jpg")
	if err != nil {
		fmt.Println(err)
	}
	dc.DrawImage(back, 0, 0)
	im := dc.Image().(*image.RGBA)
	nim := im.SubImage(image.Rect(0, 0, 1000, 100))
	f, err := os.Create("out.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = imgfactory.WriteTo(nim, f)
	_ = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
