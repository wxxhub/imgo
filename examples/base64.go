package main

import (
	"github.com/wxxhub/imgo"
)

func main() {
	base64Img := imgo.Load("gopher.png").ToBase64()
	imgo.Load(base64Img).Save("out.png")
}
