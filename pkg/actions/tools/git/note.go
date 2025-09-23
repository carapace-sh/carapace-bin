package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionNotes completes notes
//
//	e37823c1 (d10deda4)
//	fc287455 (53797853)
func ActionNotes() carapace.Action {
	return carapace.ActionExecCommand("git", "notes", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, line[41:49], line[:8])
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Note)
	}).Tag("notes").UidF(Uid("note")).QueryF(Uid("note")) // TODO ref?
}

// ActionNotesMergeStrategies completes notes merge strategies
//
//	ours (favor the local version)
//	theirs (favor the remote version)
func ActionNotesMergeStrategies() carapace.Action {
	return carapace.ActionValuesDescribed(
		"manual", "check out conflicting notes in a special work tree",
		"ours", "favor the local version",
		"theirs", "favor the remote version",
		"union", "concatenate the local and remote versions",
		"cat_sort_uniq", "similar to \"union\", but also sorts removes duplicate lines",
	)
}
