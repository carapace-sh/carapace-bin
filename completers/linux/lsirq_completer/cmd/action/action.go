package action

import (
	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"IRQ", "interrupts",
		"NAME", "name",
		"TOTAL", "total count",
	)
}
