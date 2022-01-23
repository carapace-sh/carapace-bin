package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func actionContainerExecCommand(cmd *cobra.Command, service string, index string, command string, args ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionExecCommand("docker-compose", append([]string{"exec", "--no-TTY", "--index", index, service, command}, args...)...)(func(output []byte) carapace.Action {
			return f(output)
		})
	}
}

// ActionContainerUsers completes container user names
//   root (0)
//   daemon (1)
func ActionContainerUsers(cmd *cobra.Command, service string, index string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return actionContainerExecCommand(cmd, service, index, "cat", "/etc/passwd")(func(output []byte) carapace.Action {
			users := []string{}
			for _, entry := range strings.Split(string(output), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 2 {
					user := splitted[0]
					id := splitted[2]
					if len(strings.TrimSpace(user)) > 0 {
						users = append(users, user, id)
					}
				}
			}
			return carapace.ActionValuesDescribed(users...)
		})
	})
}
