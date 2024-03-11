package golang

import "github.com/carapace-sh/carapace"

// ActionBuildmodes completes build modes
//
//	archive (Build the listed non-main packages into .a files)
//	c-archive (Build the listed main package, plus all packages it imports, into a C archive file)
func ActionBuildmodes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"archive", "Build the listed non-main packages into .a files",
		"c-archive", "Build the listed main package, plus all packages it imports, into a C archive file",
		"c-shared", "Build the listed main package, plus all packages it imports, into a C shared library",
		"default", "Listed main packages are built into executables and listed non-main packages are built into .a files",
		"shared", "Combine all the listed non-main packages into a single shared library",
		"exe", "Build the listed main packages and everything they import into executables",
		"pie", "Build the listed main packages and everything they import into position independent executables (PIE)",
		"plugin", "Build the listed main packages, plus all packages that they import, into a Go plugin",
	)
}
