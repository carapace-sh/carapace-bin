package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionRevFiles completes files of given revision
//
//	go.mod
//	go.sum
func ActionRevFiles(rev string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if rev == "" {
			rev = "@"
		}
		return actionExecJJ("file", "list", "--revision", rev)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...).MultiParts("/").StyleF(style.ForPathExt)
		})
	})
}
