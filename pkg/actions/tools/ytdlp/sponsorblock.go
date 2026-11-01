package ytdlp

import "github.com/carapace-sh/carapace"

// ActionSponsorblockCategories completes sponsorblock categories
//
//	sponsor
//	intro
func ActionSponsorblockCategories() carapace.Action {
	return carapace.ActionValues(
		"all",
		"chapter",
		"default",
		"filler",
		"interaction",
		"intro",
		"music_offtopic",
		"outro",
		"poi_highlight",
		"preview",
		"selfpromo",
		"sponsor",
	).Tag("sponsorblock categories")
}
