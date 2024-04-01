package main

import (
	"github.com/wxxhub/imgo"
	"image/color"
)

func main() {
	imgo.Canvas(500, 500, color.White).
		Insert("gopher.png", 100, 100).
		Save("out.png")
}
