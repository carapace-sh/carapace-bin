package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var graftCmd = &cobra.Command{
	Use:     "graft",
	Short:   "copy changes from other branches onto the current branch",
	GroupID: groups[group_change_manipulation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(graftCmd).Standalone()

	graftCmd.Flags().Bool("abort", false, "abort interrupted graft")
	graftCmd.Flags().String("base", "", "base revision when doing the graft merge (ADVANCED)")
	graftCmd.Flags().BoolP("continue", "c", false, "resume interrupted graft")
	graftCmd.Flags().BoolP("currentdate", "D", false, "record the current date as commit date")
	graftCmd.Flags().BoolP("currentuser", "U", false, "record the current user as committer")
	graftCmd.Flags().StringP("date", "d", "", "record the specified date as commit date")
	graftCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	graftCmd.Flags().BoolP("edit", "e", false, "invoke editor on commit messages")
	graftCmd.Flags().BoolP("force", "f", false, "force graft")
	graftCmd.Flags().Bool("log", false, "append graft info to log message")
	graftCmd.Flags().Bool("no-commit", false, "don't commit, just apply the changes in working directory")
	graftCmd.Flags().StringArrayP("rev", "r", nil, "revisions to graft")
	graftCmd.Flags().Bool("stop", false, "stop interrupted graft")
	graftCmd.Flags().String("to", "", "graft to this destination, in-memory (EXPERIMENTAL)")
	graftCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	graftCmd.Flags().StringP("user", "u", "", "record the specified user as committer")
	rootCmd.AddCommand(graftCmd)

	carapace.Gen(graftCmd).FlagCompletion(carapace.ActionMap{
		"base": action.ActionRevisions(),
		"rev":  action.ActionRevisions(),
		"tool": action.ActionMergeTools(),
	})

	carapace.Gen(graftCmd).PositionalAnyCompletion(
		action.ActionRevisions(),
	)
}
