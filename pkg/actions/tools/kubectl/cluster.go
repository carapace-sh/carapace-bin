package kubectl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionClusters completes clusters
//
//	minikube
//	another
func ActionClusters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "config", "get-clusters")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[1 : len(lines)-1]...)
		})
	})
}

// ActionUsers completes users
//
//	user1
//	user2
func ActionUsers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "config", "get-users")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[1 : len(lines)-1]...)
		})
	})
}
