package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
)

func ActionColors() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			color.ActionHexColors(),
			color.ActionXtermColorNames(),
		).Invoke(c).Merge().ToA()
	})
}
