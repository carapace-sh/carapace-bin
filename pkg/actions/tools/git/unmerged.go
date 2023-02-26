package git

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionUnmergedFiles completes unmerged files
//
//	file1.txt
//	dirA/file2.txt
func ActionUnmergedFiles() carapace.Action {
	return carapace.ActionExecCommand("git", "ls-files", "--unmerged")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				vals = append(vals, fields[len(fields)-1])
			}
		}
		return carapace.ActionValues(vals...).StyleF(style.ForPath)
	})
}
