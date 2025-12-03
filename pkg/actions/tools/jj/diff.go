package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionRevDiffs completes changes between revisions
// Accepts up to two revisions
//
//	0: compare current workspace to parent revision
//	1: compare current workspace to given revision
//	2: compare first revision to second revision
func ActionRevDiffs(revisions ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var from, to string
		switch len(revisions) {
		case 0:
			from = "@"
			to = "@-"
		case 1:
			from = revisions[0]
			to = revisions[0] + "-"
		case 2:
			from = revisions[0]
			to = revisions[1]
		default:
			return carapace.ActionMessage("invalid amount of args [ActionRevChanges]")
		}

		return carapace.ActionExecCommand("jj", "--color", "never", "diff", "--summary", "--from", from, "--to", to)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if splitted := strings.SplitN(line, " ", 2); splitted != nil {
					vals = append(vals, splitted[1], splitted[0])
				}
			}
			a := carapace.ActionValuesDescribed(vals...)
			if len(revisions) > 1 {
				a = a.MultiParts("/")
			}
			return a.StyleF(style.ForPathExt)
		})
	})
}

// ActionRevChanges completes changes made in given revisions
func ActionRevChanges(revisions ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"log", "--summary", "--no-graph", "--template", ""}
		for _, revision := range revisions {
			args = append(args, "--revisions", revision)
		}
		return carapace.ActionExecCommand("jj", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if splitted := strings.SplitN(line, " ", 2); splitted != nil {
					vals = append(vals, splitted[1], splitted[0]) // TODO state isn't always correct (will be overwritten)
				}
			}
			return carapace.ActionValuesDescribed(vals...).MultiParts("/").StyleF(style.ForPathExt)
		})
	})
}
