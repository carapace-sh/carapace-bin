package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:     "log",
	Short:   "show revision history of entire repository or files",
	Aliases: []string{"history"},
	GroupID: groups[group_change_navigation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringArrayP("bookmark", "B", nil, "show changesets within the given bookmark")
	logCmd.Flags().StringArrayP("branch", "b", nil, "show changesets within the given named branch")
	logCmd.Flags().BoolP("copies", "C", false, "show copied files")
	logCmd.Flags().StringP("date", "d", "", "show revisions matching date spec")
	logCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	logCmd.Flags().BoolP("follow", "f", false, "follow changeset history, or file history across copies and renames")
	logCmd.Flags().Bool("follow-first", false, "only follow the first parent of merge changesets (DEPRECATED)")
	logCmd.Flags().BoolP("git", "g", false, "use git extended diff format")
	logCmd.Flags().BoolP("graph", "G", false, "show the revision DAG")
	logCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	logCmd.Flags().StringArrayP("keyword", "k", nil, "do case-insensitive search for a given text")
	logCmd.Flags().StringP("limit", "l", "", "limit number of changes displayed")
	logCmd.Flags().StringArrayP("line-range", "L", nil, "follow line range of specified file (EXPERIMENTAL)")
	logCmd.Flags().BoolP("no-merges", "M", false, "do not show merges")
	logCmd.Flags().String("only-branch", "", "show only changesets within the given named branch (DEPRECATED)")
	logCmd.Flags().BoolP("only-merges", "m", false, "show only merges (DEPRECATED) (use -r \"merge()\" instead)")
	logCmd.Flags().BoolP("patch", "p", false, "show patch")
	logCmd.Flags().StringArrayP("prune", "P", nil, "do not display revision or any of its ancestors")
	logCmd.Flags().Bool("removed", false, "include revisions where files were removed")
	logCmd.Flags().StringArrayP("rev", "r", nil, "revisions to select or follow from")
	logCmd.Flags().Bool("stat", false, "output diffstat-style summary of changes")
	logCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	logCmd.Flags().StringP("template", "T", "", "display with template")
	logCmd.Flags().StringArrayP("user", "u", nil, "revisions committed by user")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"bookmark":    hg.ActionBookmarks(),
		"branch":      hg.ActionBranches(),
		"exclude":     carapace.ActionFiles(),
		"include":     carapace.ActionFiles(),
		"only-branch": hg.ActionBranches(),
		"prune":       hg.ActionRevisions(),
		"rev":         hg.ActionRevisions(),
		"user":        hg.ActionUsers(),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
