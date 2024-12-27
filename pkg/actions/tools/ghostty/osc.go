package ghostty

import "github.com/carapace-sh/carapace"

// ActionOscColorReportFormats completes osc color report formats
//
//	none (OSC 4/10/11 queries receive no reply)
//	8-bit (Color components are return unscaled)
func ActionOscColorReportFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "OSC 4/10/11 queries receive no reply",
		"8-bit", "Color components are return unscaled",
		"16-bit", "Color components are returned scaled",
	)
}
