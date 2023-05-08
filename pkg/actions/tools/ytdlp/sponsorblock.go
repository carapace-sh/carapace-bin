package ytdlp

import "github.com/rsteube/carapace"

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
	)
}
