package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
)

func ActionColors() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.CallbackValue, "#") {
			return color.ActionHexColors()
		}
		return color.Action256Colors()
	})
}
