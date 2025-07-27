package goreleaser

import "github.com/carapace-sh/carapace"

// ActionLanguages completes languages
//
//	bun
//	deno
func ActionLanguages() carapace.Action {
	return carapace.ActionValues(
		"bun",
		"deno",
		"go",
		"poetry",
		"rust",
		"uv",
		"zig",
	).Tag("languages")
}
