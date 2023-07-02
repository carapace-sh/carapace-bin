package fs

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionArFileContents completes contents of given ar file
//
//	fileA
//	fileB
func ActionArFileContents(file string) carapace.Action {
	return carapace.ActionExecCommand("ar", "-t", file)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...).StyleF(style.ForPathExt)
	})
}
