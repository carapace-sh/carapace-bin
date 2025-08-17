package docker

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"strings"
)

// ActionVolumes completes volume names
//
//	carapace-bin_go (local)
//	d35b4ebbab7bd0e9c155dbc9c75361150658ca525793af1d8fcbf97d058b905b (local)
func ActionVolumes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "volume", "ls", "--format", "{{.Name}}\n{{.Driver}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(vals[:len(vals)-1]...)
		}).Style(styles.Docker.Volume)
	}).Tag("volumes")
}
