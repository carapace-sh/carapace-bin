package action

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionSearchFormula() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("search needs at least 2 characters")
		}

		return carapace.ActionExecCommand("brew", "search", "--formula", fmt.Sprintf("/^%v/", c.CallbackValue))(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines {
				if !strings.HasPrefix(line, "==>") {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

// TODO cask
