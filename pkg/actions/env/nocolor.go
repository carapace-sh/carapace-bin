package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	knownVariables["nocolor"] = func() variables {
		return variables{
			Variables: map[string]string{
				"NO_COLOR": "disable colors in supported commands",
			},
			VariableCompletion: map[string]carapace.Action{
				"NO_COLOR": carapace.ActionStyledValuesDescribed(
					"0", "show colors", style.Carapace.KeywordNegative,
					"1", "do not show colors", style.Carapace.KeywordPositive,
				),
			},
		}
	}
}
