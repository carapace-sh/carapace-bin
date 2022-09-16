package mitmproxy

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
)

// ActionModes completes modes
//
//	regular
//	reverse
func ActionModes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			carapace.ActionValues("regular", "transparent", "socks5"),
			carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues("reverse", "upstream").Invoke(c).Suffix(":").ToA()
				case 1:
					return carapace.ActionValues("http", "https").Invoke(c).Suffix("://").ToA()
				case 2:
					return carapace.ActionValues()
				case 3:
					return net.ActionPorts()
				default:
					return carapace.ActionValues()
				}
			}),
		).ToA()
	})
}
