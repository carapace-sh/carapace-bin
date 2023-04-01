package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionBool() carapace.Action {
	return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
}
