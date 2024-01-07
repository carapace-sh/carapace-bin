package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/rsync"
)

func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasSuffix(c.Value, "%") {
			rsync.ActionFormats().Prefix(c.Value).NoSpace()
		}
		return carapace.ActionValues()
	})
}
