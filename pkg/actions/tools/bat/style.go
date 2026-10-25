package bat

import "github.com/carapace-sh/carapace"

// ActionStyles completes styles
//
//	auto (same as 'default', unless the output is piped)
//	changes (show Git modification markers)
func ActionStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"auto", "same as 'default', unless the output is piped",
		"changes", "show Git modification markers",
		"default", "enables recommended style components (default)",
		"full", "enables all available components",
		"grid", "vertical/horizontal lines to separate side bar and the header from the content",
		"header", "alias for 'header-filename'",
		"header-filename", "show filenames before the content",
		"header-filesize", "show file sizes before the content",
		"numbers", "show line numbers in the side bar",
		"plain", "disables all available components",
		"rule", "horizontal lines to delimit files",
		"snip", "draw separation lines between distinct line ranges",
	)
}
