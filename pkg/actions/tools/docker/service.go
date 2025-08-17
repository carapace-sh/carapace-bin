package docker

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"strings"
)

// ActionServices completes services
//
//	funny_robinson (alpine:latest replicated 0/1)
//	vigilant_mccarthy (bash:latest replicated 0/1)
func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "service", "ls", "--format", "{{.Name}}\n{{.Image}} {{.Mode}} {{.Replicas}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}).Style(styles.Docker.Service)
	}).Tag("services")
}
