package condition

import (
	"github.com/rsteube/carapace"
)

type Condition func(c carapace.Context) bool

// Of combines different conditions.
func Of(conditions ...Condition) Condition {
	return func(c carapace.Context) bool {
		for _, condition := range conditions {
			if !condition(c) {
				return false
			}
		}
		return true
	}
}
