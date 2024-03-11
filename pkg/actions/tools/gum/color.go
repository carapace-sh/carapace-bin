package gum

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
)

// ActionColors completes colors
//
//	121 (LightGreen)
//	#00afff (DeepSkyBlue1)
func ActionColors() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.Value, "#") {
			return color.ActionHexColors()
		}
		return color.Action256Colors()
	})
}
