package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionRevisions completes revisions, branches, bookmarks and tags
func ActionRevisions() carapace.Action {
	return carapace.ActionExecCommand("hg", "debugnamecomplete")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...)
	})
}

// ActionBookmarks completes bookmarks
func ActionBookmarks() carapace.Action {
	return carapace.ActionExecCommand("hg", "bookmarks", "-T", "{bookmark}\\t{rev}\\n")(func(output []byte) carapace.Action {
		return described(output)
	})
}

// ActionBranches completes branches
func ActionBranches() carapace.Action {
	return carapace.ActionExecCommand("hg", "branches", "-T", "{branch}\\t{rev}\\n")(func(output []byte) carapace.Action {
		return described(output)
	})
}

// ActionTags completes tags
func ActionTags() carapace.Action {
	return carapace.ActionExecCommand("hg", "tags", "-T", "{tag}\\t{rev}\\n")(func(output []byte) carapace.Action {
		return described(output)
	})
}

// ActionPaths completes remote repository path aliases
func ActionPaths() carapace.Action {
	return carapace.ActionExecCommand("hg", "paths", "-q")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...)
	})
}

// ActionShelves completes shelved changes
func ActionShelves() carapace.Action {
	return carapace.ActionExecCommand("hg", "shelve", "--list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if name, _, found := strings.Cut(line, " "); found {
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionTrackedFiles completes tracked files
func ActionTrackedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "files", "-0")(func(output []byte) carapace.Action {
		return carapace.ActionValues(nullSeparated(output)...)
	})
}

// ActionUntrackedFiles completes untracked files
func ActionUntrackedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "status", "-un")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...)
	})
}

// ActionChangedFiles completes files with uncommitted changes
func ActionChangedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "status", "-qn")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...)
	})
}

// ActionConfigKeys completes configuration option names
func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("hg", "config")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if key, value, found := strings.Cut(line, "="); found {
				vals = append(vals, key, value)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionTopics completes help topics and commands
func ActionTopics() carapace.Action {
	return carapace.ActionExecCommand("hg", "help")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if !strings.HasPrefix(line, " ") {
				continue
			}
			if name, desc, found := strings.Cut(strings.TrimSpace(line), "  "); found {
				vals = append(vals, name, strings.TrimSpace(desc))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionUsers completes user names from the commit history
func ActionUsers() carapace.Action {
	return carapace.ActionExecCommand("hg", "log", "-T", "{user}\\n")(func(output []byte) carapace.Action {
		seen := map[string]bool{}
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if !seen[line] {
				seen[line] = true
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionMergeTools completes merge tool names
func ActionMergeTools() carapace.Action {
	return carapace.ActionExecCommand("hg", "config", "merge-tools")(func(output []byte) carapace.Action {
		seen := map[string]bool{}
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if _, name, found := strings.Cut(line, "merge-tools."); found {
				if name, _, found := strings.Cut(name, "."); found && !seen[name] {
					seen[name] = true
					vals = append(vals, name)
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionUnresolvedFiles completes files with unresolved merge conflicts
func ActionUnresolvedFiles() carapace.Action {
	return carapace.ActionExecCommand("hg", "resolve", "--list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if state, path, found := strings.Cut(line, " "); found && state == "U" {
				vals = append(vals, path)
			}
		}
		return carapace.ActionValues(vals...)
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

func nullSeparated(output []byte) []string {
	vals := make([]string, 0)
	for line := range strings.SplitSeq(strings.TrimSuffix(string(output), "\x00"), "\x00") {
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
