package hg

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionRevisions completes revisions, branches, bookmarks and tags
//
//	default
//	tip
//	v1.0
func ActionRevisions() carapace.Action {
	return carapace.ActionExecCommand("hg", "debugnamecomplete")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...).Tag("revisions").UidF(Uid("revision")).QueryF(Uid("revision"))
	})
}

// ActionBookmarks completes bookmarks
//
//	bm1 (3)
//	bm2 (5)
func ActionBookmarks() carapace.Action {
	return carapace.ActionExecCommand("hg", "bookmarks", "-T", "{bookmark}\\t{rev}\\n")(func(output []byte) carapace.Action {
		return described(output).Tag("bookmarks").UidF(Uid("bookmark")).QueryF(Uid("bookmark"))
	})
}

// ActionBranches completes branches
//
//	default (3)
//	feature (5)
func ActionBranches() carapace.Action {
	return carapace.ActionExecCommand("hg", "branches", "-T", "{branch}\\t{rev}\\n")(func(output []byte) carapace.Action {
		return described(output).Tag("branches").UidF(Uid("branch")).QueryF(Uid("branch"))
	})
}

// ActionTags completes tags
//
//	tip (5)
//	v1.0 (0)
func ActionTags() carapace.Action {
	return carapace.ActionExecCommand("hg", "tags", "-T", "{tag}\\t{rev}\\n")(func(output []byte) carapace.Action {
		return described(output).Tag("tags").UidF(Uid("tag")).QueryF(Uid("tag"))
	})
}

// ActionShelves completes shelved changes
//
//	myshelf
//	another
func ActionShelves() carapace.Action {
	return carapace.ActionExecCommand("hg", "shelve", "--list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if name, _, found := strings.Cut(line, " "); found {
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...).Tag("shelves").UidF(Uid("shelf")).QueryF(Uid("shelf"))
	})
}

func lines(output []byte) []string {
	vals := make([]string, 0)
	for line := range strings.SplitSeq(strings.TrimSuffix(string(output), "\n"), "\n") {
		if line != "" {
			vals = append(vals, line)
		}
	}
	return vals
}

func described(output []byte) carapace.Action {
	vals := make([]string, 0)
	for _, line := range lines(output) {
		vals = append(vals, strings.SplitN(line, "\t", 2)...)
	}
	return carapace.ActionValuesDescribed(vals...)
}
