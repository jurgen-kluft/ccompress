package main

import (
	"github.com/jurgen-kluft/ccode"
	"github.com/jurgen-kluft/ccompress/package"
)

func main() {
	ccode.Generate(ccompress.GetPackage())
}
