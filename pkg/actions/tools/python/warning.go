package python

import "github.com/carapace-sh/carapace"

// ActionWarnings completes warnings
//
//	default (Warn once per call location)
//	error (Convert to exceptions)
func ActionWarnings() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "Warn every time",
		"default", "Warn once per call location",
		"error", "Convert to exceptions",
		"ignore", "Never warn",
		"module", "Warn once per calling module",
		"once", "Warn once per Python process",
	)
}

// ActionWarningControls completes warning controls
//
//	action:message:category:module:lineno
func ActionWarningControls() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		// TODO The full form of argument is: action:message:category:module:lineno
		switch len(c.Parts) {
		case 0:
			return ActionWarnings()
		case 1:
			return carapace.ActionValues().Usage("message")
		case 2:
			return carapace.ActionValues().Usage("category")
		case 3:
			return carapace.ActionValues().Usage("module")
		case 4:
			return carapace.ActionValues().Usage("lineno")
		default:
			return carapace.ActionValues()
		}
	})
}
