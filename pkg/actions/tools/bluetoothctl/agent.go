package bluetoothctl

import "github.com/carapace-sh/carapace"

// ActionAgentCapabilities completes agent capabilities
//
//	NoInputNoOutput
//	KeyboardOnly
func ActionAgentCapabilities() carapace.Action {
	return carapace.ActionValues(
		"DisplayOnly",
		"DisplayYesNo",
		"KeyboardDisplay",
		"KeyboardOnly",
		"NoInputNoOutput",
	).Tag("bluetooth agent capabilities")
}
