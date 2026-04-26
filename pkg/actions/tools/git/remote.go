package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionRemotes completes remote names
//
//	origin (git@github.com:some/origin.git)
//	upstream (git@github.com:some/upstream.git)
func ActionRemotes() carapace.Action {
	return carapace.ActionExecCommand("git", "remote", "--verbose")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if !strings.HasSuffix(line, " (push)") {
				continue
			}
			if remote, url, ok := strings.Cut(line, "\t"); ok {
				vals = append(vals, remote, strings.TrimSuffix(url, " (push)"))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("remotes").UidF(Uid("remote"))
}

// ActionRemoteUrls completes remote urls
func ActionRemoteUrls(remote string) carapace.Action {
	return carapace.ActionExecCommand("git", "remote", "get-url", remote)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
