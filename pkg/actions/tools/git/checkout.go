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
//	1 (main)
//	2 (feature)
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

		// C source interprets @{-N} by walking reflog in reverse
		// (object-name.c:1298-1299), so the Nth most recent checkout
		// is the Nth from the end of our forward-sorted list.
		vals := make([]string, 0)
		for i := len(branches) - 1; i >= 0; i-- {
			vals = append(vals, fmt.Sprintf("%d", len(branches)-i), branches[i])
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("prior checkouts").UidF(func(s string, uc uid.Context) (*url.URL, error) {
		return Uid("prior-checkout")(s, uc)
	})
}