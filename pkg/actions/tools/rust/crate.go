package rust

import "github.com/carapace-sh/carapace"

// ActionCrateTypes completes crate types
//
//	bin (runnable executable)
//	cdylib (dynamic system library)
func ActionCrateTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bin", "runnable executable",
		"cdylib", "dynamic system library",
		"dylib", "dynamic rust library",
		"lib", "rust library",
		"proc-macro", "procedual macro",
		"rlib", "rust library file",
		"staticlib", "static system library",
	).Tag("crate types").Uid("rust", "crate-type")
}
