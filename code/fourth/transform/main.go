package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(4000, 2000)
	back, err := gg.LoadImage("fourth/原神.jpg")
	if err != nil {
		fmt.Println(err)
	}
	dc.Scale(0.5, 0.5)
	dc.ShearAbout(0.5, 0, 2000, 1000)
	// dc.ShearAbout(0, 0.5, 2000, 1000)
	for i := 0; i <= 9; i++ {
		dc.DrawImage(back, 2000, 1000)
		dc.RotateAbout(gg.Radians(float64(10)), 2000, 1000)
	}
	dc.SavePNG("out.png")
}
