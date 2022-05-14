package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionRemotes completes remote names
//   origin
//   upstream
func ActionRemotes() carapace.Action {
	return carapace.ActionExecCommand("git", "remote")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

// ActionRemoteUrls completes remote urls
func ActionRemoteUrls(remote string) carapace.Action {
	return carapace.ActionExecCommand("git", "remote", "get-url", remote)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
