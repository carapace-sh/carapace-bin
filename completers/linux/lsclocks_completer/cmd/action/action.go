package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"CLOCK", "symbolic name",
		"ID", "numeric id",
		"ISO_TIME", "human readable ISO time",
		"NAME", "readable name",
		"NS_OFFSET", "namespace offset",
		"REL_TIME", "human readable relative time",
		"RESOL", "human readable resolution",
		"RESOL_RAW", "resolution",
		"TIME", "numeric time",
		"TYPE", "type",
	)
}

func ActionClocks() carapace.Action {
	return carapace.ActionExecCommand("lsclocks", "--output", "CLOCK")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1:]...)
	})
}
