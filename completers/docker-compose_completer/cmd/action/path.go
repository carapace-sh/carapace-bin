package action

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionServicePath(service string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path := filepath.Dir(c.CallbackValue)
		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("docker-compose", "exec", "-T", service, "ls", `-1`, `-p`, path)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		})
	})
}
