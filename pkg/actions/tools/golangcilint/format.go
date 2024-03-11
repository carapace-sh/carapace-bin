package golangcilint

import "github.com/carapace-sh/carapace"

// ActionOutFormats completes output formats
//
//	line-number
//	json
func ActionOutFormats() carapace.Action {
	return carapace.ActionValues("colored-line-number",
		"line-number",
		"json",
		"tab",
		"checkstyle",
		"code-climate",
		"html",
		"junit-xml",
		"github-actions",
		"teamcity",
	)
}
