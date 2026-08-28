package git

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/uid"
)

// ActionPriorCheckouts completes prior checkout indices and their branch names
//
//	00 (main)
//	01 (feature)
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
		// Pad with leading zeros so @{-00}..@{-99} sorts consistently
		vals := make([]string, 0)
		for i, branch := range branches {
			vals = append(vals, fmt.Sprintf("%02d", i+1), branch)
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("prior checkouts").UidF(func(s string, uc uid.Context) (*url.URL, error) {
		return Uid("prior-checkout")(s, uc)
	})
}
