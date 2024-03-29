package bluetoothctl

import "github.com/carapace-sh/carapace"

// ActionAgentCapabilities completes agent capabilities
//
//	NoInputNoOutput
//	KeyboardOnly
func ActionAgentCapabilities() carapace.Action {
	return carapace.ActionValues(
		"NoInputNoOutput",
		"KeyboardOnly",
		"KeyboardDisplay",
		"DisplayYesNo",
		"DisplayOnly",
	).Tag("bluetooth agent capabilities")
}
