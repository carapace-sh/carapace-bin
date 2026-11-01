package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

func ActionYesNo() carapace.Action {
	return carapace.ActionValues("no", "yes").StyleF(style.ForKeyword)
}
