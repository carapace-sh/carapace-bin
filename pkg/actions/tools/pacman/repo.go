package pacman

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionRepositories completes package repositories
//
//	extra
//	multilib
func ActionRepositories() carapace.Action {
	return carapace.ActionExecCommand("pacman-conf", "--repo-list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
