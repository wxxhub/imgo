package main

import "github.com/wxxhub/imgo"

func main() {
	imgo.Load("gopher.png").
		Blur(5).
		Save("out.png")
}
