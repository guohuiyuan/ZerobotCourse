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
	dc.InvertY()
	dc.DrawImage(back, 2000, 0)
	dc.SavePNG("out.png")
}
