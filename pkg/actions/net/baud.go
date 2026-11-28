package net

import "github.com/carapace-sh/carapace"

// ActionBaudRates completes baud rates
//
//	75
//	300
func ActionBaudRates() carapace.Action {
	return carapace.ActionValues(
		"75",
		"300",
		"1200",
		"2400",
		"4800",
		"9600",
		"14400",
		"19200",
		"28800",
		"38400",
		"57600",
		"115200",
		"230400",
		"460800",
		"921600",
		"1500000",
		"2000000",
	)
}
