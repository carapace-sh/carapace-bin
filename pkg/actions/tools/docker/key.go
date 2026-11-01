package docker

import (
	"github.com/carapace-sh/carapace"
)

// ActionDetachKeys completes detach keys
//
//	ctrl-@
//	ctrl-[
func ActionDetachKeys() carapace.Action {
	return carapace.ActionValues("ctrl-@", "ctrl-[", `ctrl-\`, "ctrl-]", "ctrl-^", "ctrl-_", "ctrl-{a-z}", "{a-z}")
}
