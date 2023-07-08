package tsh

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionFormats completes formats
//
//	text
//	json
func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionStyledValues(
			"text", style.ForPathExt(".txt", c),
			"json", style.ForPathExt(".json", c),
			"yaml", style.ForPathExt(".yaml", c),
		)
	})
}
