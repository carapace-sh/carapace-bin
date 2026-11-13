package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionClusterNames completes kind cluster names
func ActionClusterNames() carapace.Action {
	return carapace.ActionExecCommand("kind", "get", "clusters")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("clusters")
}

// ActionNodes completes node names, honoring the --name flag
func ActionNodes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"get", "nodes"}
		if flag := cmd.Flag("name"); flag != nil && flag.Value.String() != "" {
			args = append(args, "--name", flag.Value.String())
		}

		return carapace.ActionExecCommand("kind", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	}).Tag("nodes")
}
