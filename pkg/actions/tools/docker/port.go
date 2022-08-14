package docker

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
)

// ActionPorts completes port(range)s and protocols
//
//	80/tcp
//	100-200/udp
func ActionPorts() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return net.ActionPorts()
				case 1:
					return net.ActionPorts()
				default:
					return carapace.ActionValues()
				}
			})
		case 1:
			return carapace.ActionValues("tcp", "udp")
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionContainerPorts completes ports of a container
//
//	80/tcp
//	8080/udb
func ActionContainerPorts(container string) carapace.Action {
	return carapace.ActionExecCommand("docker", "container", "port", container)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.Fields(line)[0])
		}
		return carapace.ActionValues(vals...)
	})
}
