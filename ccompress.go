package main

import (
	"github.com/jurgen-kluft/ccode"
	ccompress "github.com/jurgen-kluft/ccompress/package"
)

func main() {
	ccode.Init()
	ccode.Generate(ccompress.GetPackage())
}
