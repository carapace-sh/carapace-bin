package conditions

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/internal/condition"
	"github.com/rsteube/carapace/pkg/traverse"
)

// ConditionParent checks if any parent directory contains one of the given file/directory
//
//	ConditionParent(".git")
//	ConditionParent("go.mod", "go.sum")
func ConditionParent(s ...string) condition.Condition {
	return func(c carapace.Context) bool {
		_, err := traverse.Parent(s...)(carapace.NewContext())
		return err == nil
	}
}
