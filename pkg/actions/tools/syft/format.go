package syft

import "github.com/rsteube/carapace"

// ActionOutputFormats completes output formats
//
//	syft-json
//	template
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"syft-json",
		"cyclonedx-xml",
		"cyclonedx-json",
		"github-json",
		"spdx-tag-value",
		"spdx-json",
		"syft-table",
		"syft-text",
		"template",
	)
}
