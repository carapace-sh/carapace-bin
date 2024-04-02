package transmission

import (
	"github.com/carapace-sh/carapace"
)

// ActionTOS Completes a name for ToS values
func ActionTOS() carapace.Action {
	return carapace.ActionValues(
		"af11",
		"af12",
		"af13",
		"af21",
		"af22",
		"af23",
		"af31",
		"af32",
		"af33",
		"af41",
		"af42",
		"af43",
		"cs0",
		"cs1",
		"cs2",
		"cs3",
		"cs4",
		"cs5",
		"cs6",
		"cs7",
		"ef",
		"le",
		"default",
		"lowcost",
		"lowdelay",
		"reliability",
		"throughput",
	)
}
