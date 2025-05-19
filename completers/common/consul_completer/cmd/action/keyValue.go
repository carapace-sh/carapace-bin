package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("consul", "kv", "get", "--keys", c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
