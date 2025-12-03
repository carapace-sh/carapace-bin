package jj

import (
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionOperations completes operations
//
//	c9f7e4956c01 (initialize repo)
//	902f106af7d3 (check out git remote's default branch)
func ActionOperations(limit int) carapace.Action {
	return actionExecJJ("op", "log", "--no-graph", "--template", `id.short() ++ "\t" ++ description ++ "\n"`, "--limit", strconv.Itoa(limit))(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, "\t", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
