package ghostty

import (
	"github.com/carapace-sh/carapace"
)

// ActionInputTypes completes input types
//
//	raw (Send raw text as-is)
//	path (Read a filepath and send the contents)
func ActionInputTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"raw", "Send raw text as-is",
		"path", "Read a filepath and send the contents",
	).Tag("input types").Uid("ghostty", "input-type")
}
