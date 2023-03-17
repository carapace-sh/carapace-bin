package kubectl

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionContexts completes contexts
//
//	minikube
//	another
func ActionContexts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "config", "get-contexts", "-o", "name")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
