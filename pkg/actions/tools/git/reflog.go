package git

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionReflogs completes reflogs
//
//	00 (subject)
//	01 (subject)
func ActionReflogs(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if ref == "" {
			ref = "HEAD"
		}
		args := []string{"log", "-g", "--pretty=tformat:%gd\t%<(64,trunc)%gs", ref}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			limit := len(lines) - 1
			vals := make([]string, 0)
			for index, line := range lines[:len(lines)-1] {
				if _, subject, ok := strings.Cut(line, "\t"); ok {
					vals = append(vals, fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(limit-1)))+"d", index), strings.TrimSpace(subject))
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("reflogs")
}
