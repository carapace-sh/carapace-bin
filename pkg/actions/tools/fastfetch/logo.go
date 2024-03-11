package fastfetch

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLogos completes logos
//
//	Alter
//	Amazon
func ActionLogos() carapace.Action {
	return carapace.ActionExecCommand("fastfetch", "--list-logos", "autocompletion")(func(output []byte) carapace.Action {
		logos := strings.Split(strings.TrimRight(string(output), "\n"), "\n")
		return carapace.Batch(
			carapace.ActionValuesDescribed(
				"none", "Disable logo",
				"small", "Show small logo if supported",
			),
			carapace.ActionValues(logos...).Style(style.Blue).Tag("builtin logos"),
		).ToA()
	})
}
