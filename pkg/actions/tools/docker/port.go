package docker

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
)

// ActionPorts completes port(range)s and protocols
//   80/tcp
//   100-200/udp
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
