package hg

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPaths completes remote repository path aliases
//
//	default
//	mirror
func ActionPaths() carapace.Action {
	return carapace.ActionExecCommand("hg", "paths", "-q")(func(output []byte) carapace.Action {
		return carapace.ActionValues(lines(output)...).Tag("remote repositories")
	})
}

// ActionConfigKeys completes configuration option names
//
//	ui.username (bob)
//	merge-tools.araxis.args (/3 /a2 ...)
func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("hg", "config")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range lines(output) {
			if key, value, found := strings.Cut(line, "="); found {
				vals = append(vals, key, value)
			}
		}
		return carapace.ActionValuesDescribed(vals...).Tag("config")
	})
}

// ActionUsers completes user names from the commit history
//
//	bob
//	alice@example.com
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
		return carapace.ActionValues(vals...).Tag("users")
	})
}

// ActionMergeTools completes merge tool names
//
//	araxis
//	beyondcompare3
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
		return carapace.ActionValues(vals...).Tag("merge tools")
	})
}

// ActionTopics completes help topics and commands
//
//	clone (make a copy of an existing repository)
//	revisions (Specifying Revisions)
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
		return carapace.ActionValuesDescribed(vals...).Tag("topics")
	})
}
