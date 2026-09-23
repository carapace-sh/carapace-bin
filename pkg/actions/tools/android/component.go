package android

import (
	"github.com/carapace-sh/carapace"
)

// ActionComponents completes component names as package/class
func ActionComponents() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPackages().Suffix("/")
		default:
			return carapace.ActionValues()
		}
	})
}
