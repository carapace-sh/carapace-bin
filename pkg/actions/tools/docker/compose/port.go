package compose

import (
	"strconv"

	"github.com/carapace-sh/carapace"
)

// ActionPorts completes ports exposed by a service
//
//	80 (8080)
func ActionPorts(opts ServicePathOpts) carapace.Action {
	return actionConfig(opts.Files, func(c config) carapace.Action {
		s, ok := c.Services[opts.Service]
		if !ok {
			return carapace.ActionMessage("unknown service %v", opts.Service)
		}

		vals := make([]string, 0)
		for _, port := range s.Ports {
			vals = append(vals, strconv.FormatUint(uint64(port.Target), 10), port.Published)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
