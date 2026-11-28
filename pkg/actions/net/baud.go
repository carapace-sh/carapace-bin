package net

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionBaudRates completes baud rates
//
//	9600 (common)
//	115200 (common)
func ActionBaudRates() carapace.Action {
	return carapace.ActionStyledValuesDescribed(
		"75", "legacy", style.Dim,
		"300", "legacy", style.Dim,
		"1200", "legacy", style.Dim,
		"2400", "legacy", style.Dim,
		"4800", "legacy", style.Dim,
		"9600", "common", style.Default,
		"14400", "legacy", style.Dim,
		"19200", "common", style.Default,
		"28800", "legacy", style.Dim,
		"38400", "common", style.Default,
		"57600", "common", style.Default,
		"115200", "common", style.Default,
		"230400", "high speed", style.Default,
		"460800", "high speed", style.Default,
		"921600", "high speed", style.Default,
		"1500000", "hardware-dependent", style.Dim,
		"2000000", "hardware-dependent", style.Dim,
	).Tag("baud rates")
}
