package action

import (
	"os/exec"

	"github.com/rsteube/carapace"
	"strings"
)

func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("docker", "config", "ls", "--format", "{{.Name}}\nupdated {{.UpdatedAt}}").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}
	})
}
