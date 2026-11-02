package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
)

func ActionAliases() carapace.Action {
	return gh.ActionAliases()
}
