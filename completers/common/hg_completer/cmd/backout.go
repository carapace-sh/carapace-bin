package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var backoutCmd = &cobra.Command{
	Use:     "backout",
	Short:   "reverse effect of earlier changeset",
	GroupID: groups[group_change_manipulation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backoutCmd).Standalone()

	backoutCmd.Flags().Bool("commit", false, "commit if no conflicts were encountered (DEPRECATED)")
	backoutCmd.Flags().StringP("date", "d", "", "record the specified date as commit date")
	backoutCmd.Flags().BoolP("edit", "e", false, "invoke editor on commit messages")
	backoutCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	backoutCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	backoutCmd.Flags().StringP("logfile", "l", "", "read commit message from file")
	backoutCmd.Flags().Bool("merge", false, "merge with old dirstate parent after backout")
	backoutCmd.Flags().StringP("message", "m", "", "use text as commit message")
	backoutCmd.Flags().Bool("no-commit", false, "do not commit")
	backoutCmd.Flags().String("parent", "", "parent to choose when backing out merge (DEPRECATED)")
	backoutCmd.Flags().StringP("rev", "r", "", "revision to backout")
	backoutCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	backoutCmd.Flags().StringP("user", "u", "", "record the specified user as committer")
	rootCmd.AddCommand(backoutCmd)

	carapace.Gen(backoutCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"logfile": carapace.ActionFiles(),
		"parent":  hg.ActionRevisions(),
		"rev":     hg.ActionRevisions(),
		"tool":    hg.ActionMergeTools(),
	})

	carapace.Gen(backoutCmd).PositionalCompletion(
		hg.ActionRevisions(),
	)
}
