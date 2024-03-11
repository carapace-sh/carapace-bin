package bat

import "github.com/carapace-sh/carapace"

// ActionStyles completes styles
//
//	auto (same as 'default', unless the output is piped)
//	changes (show Git modification markers)
func ActionStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"default", "enables recommended style components (default)",
		"full", "enables all available components",
		"auto", "same as 'default', unless the output is piped",
		"plain", "disables all available components",
		"changes", "show Git modification markers",
		"header", "alias for 'header-filename'",
		"header-filename", "show filenames before the content",
		"header-filesize", "show file sizes before the content",
		"grid", "vertical/horizontal lines to separate side bar and the header from the content",
		"rule", "horizontal lines to delimit files",
		"numbers", "show line numbers in the side bar",
		"snip", "draw separation lines between distinct line ranges",
	)
}
