package git

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionPriorCheckouts completes prior checkout indices and their branch names
//
//	01 (main)
//	02 (feature)
func ActionPriorCheckouts() carapace.Action {
	return carapace.ActionExecCommand("git", "reflog", "--format=%gs", "HEAD")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		branches := make([]string, 0)
		for _, line := range lines {
			if !strings.HasPrefix(line, "checkout: moving from ") {
				continue
			}
			rest := strings.TrimPrefix(line, "checkout: moving from ")
			from, _, ok := strings.Cut(rest, " to ")
			if !ok {
				continue
			}
			branches = append(branches, from)
		}

		// C source walks reflog in reverse (newest first), decrementing
		// remaining until 0 (object-name.c:1243-1267). @{-1} is the from
		// branch of the most recent checkout entry.
		// Limit to 99 entries as the C source only validates @{-N} for N > 0
		// (object-name.c:1293).
		limit := min(len(branches), 99)
		vals := make([]string, 0)
		for i, branch := range branches[:limit] {
			vals = append(vals, fmt.Sprintf("%02d", i+1), branch)
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("prior checkouts").UidF(Uid("prior-checkout"))
}
