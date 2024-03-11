package zoxide

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionDirectories completes zoxide directories
//
//	/tmp/
//	/tmp/dirA/
func ActionDirectories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("zoxide", "query", "--list", c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			for index := range lines {
				lines[index] += "/"
			}
			return carapace.ActionValues(lines[:len(lines)-1]...).MultiParts("/").StyleF(style.ForPathExt)
		})
	})
}
