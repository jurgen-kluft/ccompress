package ccompress

import (
	cbase "github.com/jurgen-kluft/cbase/package"
	"github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

// GetPackage returns the package object of 'ccompress'
func GetPackage() *denv.Package {
	// Dependencies
	unittestpkg := cunittest.GetPackage()
	basepkg := cbase.GetPackage()

	// The main (ccompress) package
	mainpkg := denv.NewPackage("ccompress")
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(basepkg)

	// 'ccompress' library
	mainlib := denv.SetupDefaultCppLibProject("ccompress", "github.com\\jurgen-kluft\\ccompress")
	mainlib.Dependencies = append(mainlib.Dependencies, basepkg.GetMainLib())

	// 'ccompress' unittest project
	maintest := denv.SetupDefaultCppTestProject("ccompress_test", "github.com\\jurgen-kluft\\ccompress")
	maintest.Dependencies = append(maintest.Dependencies, unittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, basepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
