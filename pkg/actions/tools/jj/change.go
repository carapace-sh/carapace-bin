package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionChangeIds completes change ids
//
//	t (message)
//	x (another message)
func ActionChangeIds() carapace.Action {
	return actionExecJJ("log", "--no-graph", "--template", `change_id.shortest() ++ "\n" ++ description.first_line() ++ "\n"`)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
	}).Tag("change ids")
}
