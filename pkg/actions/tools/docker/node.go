package docker

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"strings"
)

// ActionNodes completes node ids
//
//	r08tybjkcyar8vdglerxxxxxx
//	r08tybjkcyar8vdgleryyyyyy
func ActionNodes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "node", "ls", "--format", "{{.ID}}\n{{.Hostname}} {{.ManagerStatus}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...).Style(styles.Docker.Node)
		})
	}).Tag("nodes")
}
