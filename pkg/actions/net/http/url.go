package http

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
)

// ActionUrls completes known hosts and ports as urls
//
//	aur.archlinux.org
//	https://github.com:80
func ActionUrls() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		prefix := ""

		if strings.HasPrefix(c.CallbackValue, "https://") {
			c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "https://")
			prefix = "https://"
		} else if strings.HasPrefix(c.CallbackValue, "http://") {
			c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "http://")
			prefix = "http://"
		}

		return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}).Invoke(c).Prefix(prefix).ToA()
	})
}
