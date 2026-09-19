package hg

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTrackedFiles completes tracked files
//
//	.hgtags
//	a.txt
func ActionTrackedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "files", "-0")(func(output []byte) carapace.Action {
		return carapace.ActionValues(nullSeparated(output)...).Tag("tracked files").UidF(Uid("tracked-file")).QueryF(Uid("tracked-file"))
	})
}

// ActionUntrackedFiles completes untracked files
//
//	b.txt
func ActionUntrackedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "status", "-un")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...).Tag("untracked files").UidF(Uid("untracked-file")).QueryF(Uid("untracked-file"))
	})
}

// ActionChangedFiles completes files with uncommitted changes
//
//	a.txt
func ActionChangedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "status", "-qn")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...).Tag("changed files").UidF(Uid("changed-file")).QueryF(Uid("changed-file"))
	})
}

// ActionUnresolvedFiles completes files with unresolved merge conflicts
//
//	c.txt
func ActionUnresolvedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "resolve", "--list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if state, path, found := strings.Cut(line, " "); found && state == "U" {
				vals = append(vals, path)
			}
		}
		return carapace.ActionValues(vals...).Tag("unresolved files").UidF(Uid("unresolved-file")).QueryF(Uid("unresolved-file"))
	})
}

func nullSeparated(output []byte) []string {
	vals := make([]string, 0)
	for line := range strings.SplitSeq(strings.TrimSuffix(string(output), "\x00"), "\x00") {
		if line != "" {
			vals = append(vals, line)
		}
	}
	return vals
}
