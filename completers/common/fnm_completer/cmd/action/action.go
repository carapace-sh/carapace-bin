package action

import "github.com/carapace-sh/carapace"

func ActionLogLevel() carapace.Action {
	return carapace.ActionValues(
		"quiet",
		"info",
		"error",
	)
}

func ActionVersionFileStrategy() carapace.Action {
	return carapace.ActionValuesDescribed(
		"local", "Use the local version of Node defined within the current directory",
		"recursive", "Use the version of Node defined within the current dirctory and all parent directories",
	)
}

func ActionResolveEngines() carapace.Action {
	return carapace.ActionValues("true", "false")
}
