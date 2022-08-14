package compose

import (
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
)

func actionContainerExecCommand(files []string, service string, index int, command string, args ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return actionExecCompose(files, append([]string{"exec", "--no-TTY", "--index", strconv.Itoa(index), service, command}, args...)...)(func(output []byte) carapace.Action {
			return f(output)
		})
	}
}

type ContainerUserOpts struct {
	Files   []string
	Service string
	Index   int
}

// ActionUsers completes users within a service container
//
//	root (0)
//	daemon (1)
func ActionUsers(opts ContainerUserOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if opts.Index == 0 {
			opts.Index = 1 // default to 1
		}

		return actionContainerExecCommand(opts.Files, opts.Service, opts.Index, "cat", "/etc/passwd")(func(output []byte) carapace.Action {
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
