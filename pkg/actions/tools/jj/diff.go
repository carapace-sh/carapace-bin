package jj

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionRevDiffs completes changes beetween revisions
// Accepts up to two revisions
//
//	0: compare current workspace to parent revision
//	1: compare current workspace to given revision
//	2: compare first revision to second revision
func ActionRevDiffs(revisions ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var from, to string
		switch len(revisions) {
		case 0:
			from = "@"
			to = "@-"
		case 1:
			from = revisions[0]
			to = revisions[0] + "-"
		case 2:
			from = revisions[0]
			to = revisions[1]
		default:
			return carapace.ActionMessage("invalid amount of args [ActionRevChanges]")
		}

		return carapace.ActionExecCommand("jj", "diff", "--summary", "--from", from, "--to", to)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if splitted := strings.SplitN(line, " ", 2); splitted != nil {
					vals = append(vals, splitted[1], splitted[0])
				}
			}
			a := carapace.ActionValuesDescribed(vals...)
			if len(revisions) > 1 {
				a = a.MultiParts("/")
			}
			return a.StyleF(style.ForPathExt)
		})
	})
}
