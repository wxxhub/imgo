# ImGo
English | [简体中文](README-CN.md)

![Github stars](https://img.shields.io/github/stars/fishtailstudio/imgo?style=social)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fishtailstudio/imgo)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/fishtailstudio/imgo)
[![GoDoc](https://godoc.org/github.com/wxxhub/imgo?status.svg)](https://pkg.go.dev/github.com/wxxhub/imgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/wxxhub/imgo)](https://goreportcard.com/report/github.com/wxxhub/imgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Introduction

`Image Golang` => `Img Go` => `ImGo` `[ˈɪmɡəʊ]`

`ImGo` is an open source Golang image handling and manipulation library.  It provides an easier and expressive way to create, edit, and compose images.

## Installation

```bash
go get -u github.com/wxxhub/imgo
```

## Documentation

[English Documentation](https://imgo.gitbook.io/en/) | [简体中文文档](https://imgo.gitbook.io/cn/)

## Usage

```go
package main

import "github.com/wxxhub/imgo"

func main() {
    imgo.Load("background.png").
        Resize(250, 350).
        Insert("gopher.png", 50, 50).
        Save("out.png")
}
```

## Maintainers

[@fishtailstudio](https://github.com/fishtailstudio)

## Contributing

Feel free to dive in! [Open an Issue](https://github.com/wxxhub/imgo/issues/new) or [submit PRs](https://github.com/wxxhub/imgo/pulls).

## Give a Star ! ⭐

If you like or are using this project, please give it a star. Thanks!

## License

[MIT](LICENSE) © fishtailstudio