package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
)

// ActionInstallUsers completes the user argument Package Manager accepts:
// a device user id, or `all` / `current`.
func ActionInstallUsers() carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed(
			"all", "all users",
			"current", "the current foreground user",
		),
		adb.ActionUsers(),
	).ToA()
}
