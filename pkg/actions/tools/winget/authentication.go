package winget

import "github.com/carapace-sh/carapace"

// ActionAuthenticationModes completes authentication modes
//
//	interactive
//	silent
func ActionAuthenticationModes() carapace.Action {
	return carapace.ActionValues(
		"interactive",
		"silent",
		"silentPreferred",
	).Tag("authentication modes")
}
