package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

type RevOption struct {
	LocalBookmarks  bool
	RemoteBookmarks bool
	Commits         int
	HeadCommits     int
	Tags            bool
}

func (o RevOption) Default() RevOption {
	o.LocalBookmarks = true
	o.RemoteBookmarks = true
	o.Commits = 100
	o.HeadCommits = 20
	o.Tags = true
	return o

}

// ActionRevs completes refs (commits, bookmarks, tags)
func ActionRevs(revOption RevOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()

		if revOption.LocalBookmarks {
			batch = append(batch, ActionLocalBookmarks())
		}

		if revOption.RemoteBookmarks {
			batch = append(batch, ActionRemoteBookmarks(""))
		}

		if revOption.Commits > 0 {
			batch = append(batch, ActionRecentCommits(revOption.Commits))
		}

		if revOption.HeadCommits > 0 {
			batch = append(batch, ActionHeadCommits(revOption.HeadCommits))
		}

		if revOption.Tags {
			batch = append(batch, ActionTags())
		}

		return batch.ToA()
	})
}

// ActionRevSets completes revision sets
func ActionRevSets(opts RevOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO very basic at the moment
		index := strings.LastIndexAny(c.Value, " .:())|+&")
		prefix := c.Value[:index+1]

		c.Value = strings.TrimPrefix(c.Value, prefix)
		return carapace.Batch(
			ActionRevs(opts),
			ActionRevSetFunctions().Suffix("("),
		).ToA().Invoke(c).Prefix(prefix).ToA().NoSpace()
	})
}

// ActionRevSetFunctions completes refset functions
//
//	parents (Same as x-)
//	children (Same as x+)
func ActionRevSetFunctions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"parents", "Same as x-",
		"children", "Same as x+",
		"ancestors", "Ancestors of x limited to the given depth",
		"descendants", "Same as x::",
		"connected", "Same as x::x",
		"all", "All visible commits in the repo",
		"none", "No commits",
		"bookmarks", "All local bookmark targets",
		"remote_bookmarks", "All remote bookmark targets across all remotes",
		"tags", "All tag targets",
		"git_refs", "All Git ref targets as of the last import",
		"git_head", "The Git HEAD target as of the last import",
		"visible_heads", "All visible heads (same as heads(all()))",
		"root", "The virtual commit that is the oldest ancestor of all other commits",
		"heads", "Commits in x that are not ancestors of other commits in x",
		"roots", "Commits in x that are not descendants of other commits in x",
		"latest", "Latest count commits in x, based on committer timestamp",
		"merges", "Merge commits",
		"description", "Commits that have a description matching the given string pattern",
		"author", "Commits with the author's name or email matching the given string pattern",
		"mine", "Commits where the author's email matches the email of the current user",
		"committer", "Commits with the committer's name or email matching the given string pattern",
		"empty", "Commits modifying no files. This also includes merges() without user modifications and root()",
		"files", "Commits modifying paths matching the given fileset expression",
		"conflicts", "Commits with conflicts",
		"present", "Same as x, but evaluated to none() if any of the commits in x doesn't exist",
		"reachable", "All commits reachable from srcs within domain",
		"mutable", "All commits that jj does not treat as immutable (same as ~immutable())",
		"immutable", "All commits that jj treats as immutable (same as (immutable_heads() | root()))",
		"diff_contains", "Commits containing diffs matching the given text pattern line by line",
		"author_date", "Commits with author dates matching the specified date pattern",
		"committer_date", "Commits with committer dates matching the specified date pattern",
		"tracked_remote_bookmarks", "All targets of tracked remote bookmarks",
		"untracked_remote_bookmarks", "All targets of untracked remote bookmarks",
		"coalesce", "Get first non-none revset from a list of revsets",
		"at_operation", "Query revisions based on historical state",
		"fork_point", "Obtain the fork point of multiple commits",
	).Tag("revset functions")
}
