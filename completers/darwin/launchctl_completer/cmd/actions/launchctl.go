package actions

import (
	"os"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionServices() carapace.Action {
	return carapace.ActionExecCommand("launchctl", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		values := make([]string, 0, len(lines)*2)
		for _, line := range lines[1:] {
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}
			values = append(values, fields[2], strings.Join(fields[:2], " "))
		}
		return carapace.ActionValuesDescribed(values...)
	})
}

func ActionDomains() carapace.Action {
	uid := strconv.Itoa(os.Getuid())
	return carapace.ActionValuesDescribed(
		"system", "system services",
		"user/"+uid, "current user's per-user services",
		"gui/"+uid, "current GUI login session services",
	)
}
