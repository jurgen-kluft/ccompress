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
	mainpkg := denv.NewPackage("github.com\\jurgen-kluft", "ccompress")
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(basepkg)

	// 'ccompress' library
	mainlib := denv.SetupCppLibProject(mainpkg, "ccompress")
	mainlib.AddDependencies(basepkg.GetMainLib()...)

	// 'ccompress' unittest project
	maintest := denv.SetupCppTestProject(mainpkg, "ccompress_test")
	maintest.AddDependencies(unittestpkg.GetMainLib()...)
	maintest.AddDependencies(basepkg.GetMainLib()...)
	maintest.AddDependency(mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
