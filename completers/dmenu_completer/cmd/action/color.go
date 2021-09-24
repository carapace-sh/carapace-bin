package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
)

func ActionColors() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			os.ActionHexColors(),
			os.ActionXtermColorNames(),
		).Invoke(c).Merge().ToA()
	})
}
