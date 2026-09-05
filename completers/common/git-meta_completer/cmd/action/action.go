package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

func ActionTarget() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"branch", "branch",
				"change-id", "change-id",
				"commit", "commit",
				"path", "path",
				"project", "project",
			).Suffix(":")
		default:
			switch c.Parts[0] {
			case "commit":
				return git.ActionRefs(git.RefOption{}.Default())
			case "branch":
				return git.ActionLocalBranches()
			case "path":
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}
	})
}
