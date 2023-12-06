package pacman

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionCountries completes mirror countries
//
//	Armenia
//	Australia
func ActionCountries() carapace.Action {
	return carapace.ActionExecCommand("pacman-mirrors", "--country-list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
