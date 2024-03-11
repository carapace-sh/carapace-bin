package ytdlp

import "github.com/carapace-sh/carapace"

// ActionSponsorblockCategories completes sponsorblock categories
//
//	sponsor
//	intro
func ActionSponsorblockCategories() carapace.Action {
	return carapace.ActionValues(
		"sponsor",
		"intro",
		"outro",
		"selfpromo",
		"preview",
		"filler",
		"interaction",
		"music_offtopic",
		"poi_highlight",
		"chapter",
		"all",
		"default",
	).Tag("sponsorblock categories")
}
