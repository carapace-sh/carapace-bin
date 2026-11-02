package action

import (
	"github.com/carapace-sh/carapace"
)

// ActionColors completes the values of the --color flag
func ActionColors() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "Always display colors",
		"auto", "Display colors if the output goes to an interactive terminal",
		"never", "Never display colors",
	).Tag("colors")
}

// ActionOutputFormats completes lint output formats
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"azure",
		"concise",
		"full",
		"github",
		"gitlab",
		"grouped",
		"json",
		"json-lines",
		"junit",
		"pylint",
		"rdjson",
		"sarif",
	).Tag("output formats")
}

// ActionPythonVersions completes supported Python versions
func ActionPythonVersions() carapace.Action {
	return carapace.ActionValues(
		"py37",
		"py38",
		"py39",
		"py310",
		"py311",
		"py312",
		"py313",
		"py314",
		"py315",
	).Tag("python versions")
}

// ActionDirections completes graph directions
func ActionDirections() carapace.Action {
	return carapace.ActionValuesDescribed(
		"dependencies", "Construct a map from module to its dependencies (i.e., the modules that it imports)",
		"dependents", "Construct a map from module to its dependents (i.e., the modules that import it)",
	).Tag("directions")
}

// ActionShells completes shell names
func ActionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nushell", "powershell", "zsh")
}
