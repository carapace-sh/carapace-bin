package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rsync"
)

func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasSuffix(c.Value, "%") {
			rsync.ActionFormats().Prefix(c.Value).NoSpace()
		}
		return carapace.ActionValues()
	})
}
