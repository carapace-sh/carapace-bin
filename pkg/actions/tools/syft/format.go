package syft

import "github.com/carapace-sh/carapace"

// ActionOutputFormats completes output formats
//
//	syft-json
//	template
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"cyclonedx-json",
		"cyclonedx-xml",
		"github-json",
		"spdx-json",
		"spdx-tag-value",
		"syft-json",
		"syft-table",
		"syft-text",
		"template",
	)
}
