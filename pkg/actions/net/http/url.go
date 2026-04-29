package http

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
)

// ActionUrls completes known hosts and ports as urls
//
//	aur.archlinux.org
//	https://github.com:80
func ActionUrls() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		prefix := ""

		if after, ok := strings.CutPrefix(c.Value, "https://"); ok {
			c.Value = after
			prefix = "https://"
		} else if after, ok := strings.CutPrefix(c.Value, "http://"); ok {
			c.Value = after
			prefix = "http://"
		}

		return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().NoSpace()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}).Invoke(c).Prefix(prefix).ToA()
	})
}
