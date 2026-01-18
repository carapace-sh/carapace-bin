package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"TYPE", "type",
		"ID", "numeric id",
		"CLOCK", "symbolic name",
		"NAME", "readable name",
		"TIME", "numeric time",
		"ISO_TIME", "human readable ISO time",
		"RESOL", "human readable resolution",
		"RESOL_RAW", "resolution",
		"REL_TIME", "human readable relative time",
		"NS_OFFSET", "namespace offset",
	)
}

func ActionClocks() carapace.Action {
	return carapace.ActionExecCommand("lsclocks", "--output", "CLOCK")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1:]...)
	})
}
