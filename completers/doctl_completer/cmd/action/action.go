package action

import "github.com/rsteube/carapace"

func ActionActionStatus() carapace.Action {
	// TODO incomplete
	return carapace.ActionValues("pending", "completed")
}

func ActionActionTypes() carapace.Action {
	// TODO incomplete
	return carapace.ActionValues("create", "destroy", "power_cycle", "power_off", "power_on", "backup", "migrate", "attach_volume")
}
