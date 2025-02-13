package golang

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTools completes tools
//
//	asm
//	buildid
func ActionTools() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("go", "tool")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				switch {
				case strings.Contains(line, "/"):
					// Use base of module for convenience.
					// Name clashes are possible but should be unlikely.
					vals = append(vals, filepath.Base(line), line)
				default:
					vals = append(vals, line, "")
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
