package net

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionBaudRates completes baud rates
//
//	9600
//	115200
func ActionBaudRates() carapace.Action {
	return carapace.ActionStyledValues(
		"75", style.Dim,
		"300", style.Dim,
		"1200", style.Dim,
		"2400", style.Dim,
		"4800", style.Dim,
		"9600", style.Default,
		"14400", style.Dim,
		"19200", style.Default,
		"28800", style.Dim,
		"38400", style.Default,
		"57600", style.Default,
		"115200", style.Default,
		"230400", style.Default,
		"460800", style.Default,
		"921600", style.Default,
		"1500000", style.Dim,
		"2000000", style.Dim,
	).Tag("baud rates")
}
