package action

import (
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionCodesignIdentities() carapace.Action {
	if runtime.GOOS != "darwin" {
		return carapace.ActionValues()
	}
	return carapace.ActionExecCommand("security", "find-identity", "-v", "-p", "codesigning")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if i := strings.Index(line, "\""); i != -1 {
				if j := strings.LastIndex(line, "\""); j > i {
					vals = append(vals, line[i+1:j])
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
