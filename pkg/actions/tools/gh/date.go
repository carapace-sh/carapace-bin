package gh

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionDateFields completes date fields
//
//	<=2023-01-01
//	2022-01-01...2023-01-01
func ActionDateFields() carapace.Action {
	return carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			operators := []string{">=", ">", "<=", "<"}
			prefix := ""
			for _, operator := range operators {
				if strings.HasPrefix(c.Value, operator) {
					c.Value = strings.TrimPrefix(c.Value, operator)
					prefix = operator
					break
				}
			}
			return carapace.Batch(
				carapace.ActionValues(">", ">=", "<", "<=").Style(style.Blue).NoSpace('>', '=', '<'),
				time.ActionDate().Invoke(c).Prefix(prefix).ToA().NoSpace(),
			).ToA()

		case 1:
			return time.ActionDate()

		default:
			return carapace.ActionValues()
		}
	})
}
