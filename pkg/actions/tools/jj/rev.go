package jj

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	jjlex "github.com/carapace-sh/carapace-jjlex"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/pelletier/go-toml"
)

type RevOption struct {
	LocalBookmarks  bool
	RemoteBookmarks bool
	Commits         int
	HeadCommits     int
	Tags            bool
	ChangeIds       bool
}

func (o RevOption) Default() RevOption {
	o.LocalBookmarks = true
	o.RemoteBookmarks = true
	o.Commits = 100
	o.HeadCommits = 1
	o.Tags = true
	o.ChangeIds = true
	return o

}

// ActionRevs completes revs (commits, bookmarks, tags)
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

		if revOption.ChangeIds {
			batch = append(batch, ActionChangeIds())
		}

		return batch.ToA()
	})
}

// ActionRevsets completes revsets
func ActionRevsets(opts RevOption) carapace.Action { // TODO remove opts
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()       // default prefix
		revsetBatch := carapace.Batch() // attachedRevset prefix

		ctx := jjlex.Split(c.Value)
		switch ctx.Type {
		case jjlex.CompletionTypeOperator:
			attached := strings.HasSuffix(strings.TrimSuffix(ctx.FullInput, ctx.Prefix), " ")
			batch = append(batch, ActionRevsetOperators(attached))
			revsetBatch = append(revsetBatch,
				ActionAncestors(ctx.AttachedRevset).
					Suppress("doesn't exist"). // revset might be an incomplete bookmark or similar that contains `-`
					Unless(ctx.AttachedRevset == "" || (!strings.HasSuffix(ctx.AttachedRevset, "-") && !strings.HasSuffix(ctx.AttachedRevset, "+"))),
				ActionDescendants(ctx.AttachedRevset).
					Suppress("doesn't exist"). // revset might be an incomplete bookmark or similar that contains `+`
					Unless(ctx.AttachedRevset == "" || (!strings.HasSuffix(ctx.AttachedRevset, "+") && !strings.HasSuffix(ctx.AttachedRevset, "-"))),
			)
		case jjlex.CompletionTypeFunctionArg:
			// TODO complete corresponding type (e.g. lexer should return revision)
			batch = append(batch,
				ActionRevs(RevOption{}.Default()),
				ActionRevsetFunctions().Suffix("("),
			)
		case jjlex.CompletionTypeRevision:
			fullPrefix := strings.TrimSuffix(ctx.FullInput, ctx.Prefix)
			attached := strings.HasSuffix(fullPrefix, " ")
			batch = append(batch,
				ActionRevs(RevOption{}.Default()),
				ActionRevsetFunctions().Suffix("("),
				ActionRevsetAliases().Style(style.Dim),
			)
			switch {
			case strings.HasSuffix(fullPrefix, ".."), strings.HasSuffix(fullPrefix, "::"):
				batch = append(batch, ActionRevsetOperators(attached).Filter("..", "::").Prefix(ctx.Prefix)) // `revA....revB` is not allowed
			default:
				revsetBatch = append(revsetBatch,
					ActionAncestors(ctx.AttachedRevset).
						Suppress("doesn't exist"). // revset might be an incomplete bookmark or similar that contains `-`
						Unless(ctx.AttachedRevset == "" || (!strings.HasSuffix(ctx.AttachedRevset, "-") && !strings.HasSuffix(ctx.AttachedRevset, "+"))),
					ActionDescendants(ctx.AttachedRevset).
						Suppress("doesn't exist"). // revset might be an incomplete bookmark or similar that contains `+`
						Unless(ctx.AttachedRevset == "" || (!strings.HasSuffix(ctx.AttachedRevset, "+") && !strings.HasSuffix(ctx.AttachedRevset, "-"))),
				)
			}
		}
		c.Value = ctx.Prefix
		return carapace.Batch(
			batch.ToA().Prefix(strings.TrimSuffix(ctx.FullInput, ctx.Prefix)),
			revsetBatch.ToA().Prefix(strings.TrimSuffix(ctx.FullInput, ctx.AttachedRevset)),
		).ToA().Invoke(c).ToA().NoSpace()
	})
}

// ActionRevsetFunctions completes revset functions
//
//	parents (Same as x-)
//	children (Same as x+)
func ActionRevsetFunctions() carapace.Action {
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
	).Tag("revset functions").Uid("jj", "revset-function")
}

// ActionRevsetOperators completes revset operators
//
//	\- (x-`: Parents of x, can be empty)
//	+ (x+`: Children of x, can be empty)
func ActionRevsetOperators(attached bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()

		switch {
		case attached:
			batch = append(batch, carapace.ActionValuesDescribed(
				"-", "`x-`: Parents of x, can be empty",
				"+", "`x+`: Children of x, can be empty",
				":", "`p:x`: String/date pattern or pattern alias named p",
				"::", "`x::`: Descendants of x, including the commits in x itself",
				"..", "`x..`: Revisions that are not ancestors of x",
			))
		default:
			batch = append(batch, carapace.ActionValuesDescribed(
				"::", "`::x`: Ancestors of x, including the commits in x itself",
				"..", "`..x`: Ancestors of x, including the commits in x itself",
				"~", "`~x`: Revisions that are not in x",
			))
		}

		batch = append(batch, carapace.ActionValuesDescribed(
			"&", "`x & y`: Revisions that are in both x and y",
			"~", "`x ~ y`: Revisions that are in x but not in y",
			"|", "`x | y`: Revisions that are in either x or y (or both)",
		))

		return batch.ToA().Uid("jj", "revset-operator", "attached", strconv.FormatBool(attached))
	}).Tag("revset operators")
}

// ActionRevsetAliases completes revset aliases
//
//	HEAD (@-)
//	trunk() (main@origin)
func ActionRevsetAliases() carapace.Action {
	return actionExecJJ("config", "list", "revset-aliases")(func(output []byte) carapace.Action {
		var aliases struct {
			RevsetAliases map[string]string `toml:"revset-aliases"`
		}
		if err := toml.Unmarshal(output, &aliases); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(aliases.RevsetAliases))
		for name, alias := range aliases.RevsetAliases {
			// TODO name can contain parameters which need to be handled (conditionally? see ActionRevSets)
			// 'HEAD' = '@-'
			// 'user()' = 'user("me@example.org")'
			// 'user(x)' = 'author(x) | committer(x)'
			// 'grep:x' = 'description(regex:x)'

			vals = append(vals, name, alias)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("revset aliases").UidF(Uid("revset")) // TODO revset-alias (parameter will be an issue)
}

// ActionAncestors completes ancestors
//
//	\- (message)
//	-- (message)
func ActionAncestors(revset string) carapace.Action { // TODO revset unused (clashes with c.value below)
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if revset == "" {
			revset = "@"
		}
		return actionExecJJ("log", "--no-graph", "--template", `description.first_line() ++ "\n"`, "--revisions", fmt.Sprintf("first_ancestors(%v)", revset), "--limit", "20")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for index, line := range lines[:len(lines)-1] {
				if index == 0 {
					continue
				}
				vals = append(vals, strings.Repeat("-", index), line)
			}
			return carapace.ActionValuesDescribed(vals...).Prefix(revset)
		})
	}).Tag("ancestors").UidF(Uid("revset"))
}

// ActionDescendants completes descendants
//
//	\+ (message)
//	++ (message)
func ActionDescendants(revset string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if revset == "" {
			revset = "@"
		}
		batch := carapace.Batch()
		for i := range 20 {
			// TODO missing the first descendant
			batch = append(batch, actionExecJJE("show", "--template", `description.first_line()`, revset+strings.Repeat("+", i+1))(func(output []byte, err error) carapace.Action {
				if err != nil {
					if exitErr, ok := err.(*exec.ExitError); ok {
						switch {
						case strings.Contains(string(exitErr.Stderr), "resolved to more than one revision"):
							return carapace.ActionValuesDescribed(strings.Repeat("+", i+1), "resolves to more than one revision").Prefix(c.Value).Style(style.Carapace.KeywordAmbiguous)
						case strings.Contains(string(exitErr.Stderr), "didn't resolve to any revisions"):
							return carapace.ActionValues()
						default:
							if firstLine := strings.SplitN(string(exitErr.Stderr), "\n", 2)[0]; strings.TrimSpace(firstLine) != "" {
								err = errors.New(firstLine)
							}
						}
					}
					return carapace.ActionMessage(err.Error())
				}
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValuesDescribed(strings.Repeat("+", i+1), lines[0]).Prefix(revset)
			}).Invoke(c).ToA())
		}
		return batch.ToA()
	}).Tag("descendants").UidF(Uid("revset"))
}
