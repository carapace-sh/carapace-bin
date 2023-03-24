package git

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionNotes completes notes

func ActionNotes() carapace.Action {
	return carapace.ActionExecCommand("git", "log", "--pretty=format:%h\n%<(64,trunc)%s", "--show-notes")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValuesDescribed(lines...).Style(styles.Git.Note)
	}).Tag("notes")
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
