package flatpak

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionArches completes architectures
//
//	i386
//	x86_64
func ActionArches() carapace.Action {
	return carapace.ActionExecCommand("flatpak", "--supported-arches")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
