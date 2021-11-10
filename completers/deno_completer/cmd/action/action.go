package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
)

func ActionHostPort() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValues() // TODO local ips
		case 1:
			return net.ActionPorts()
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionLintRules() carapace.Action {
	return carapace.ActionExecCommand("deno", "lint", "--rules")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if strings.HasPrefix(line, " - ") {
				vals = append(vals, strings.TrimPrefix(line, " - "))
			}
		}
		return carapace.ActionValues(vals...)
	})
}
