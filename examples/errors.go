package main

import (
	"fmt"
	"github.com/wxxhub/imgo"
)

func main() {
	err := imgo.Load("gopher.jpg").Save("out.png").Error
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("success")
	}
}
