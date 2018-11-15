package main

import (
	"github.com/jurgen-kluft/xcode"
	"github.com/jurgen-kluft/xcompress/package"
)

func main() {
	xcode.Init()
	xcode.Generate(xcompress.GetPackage())
}
