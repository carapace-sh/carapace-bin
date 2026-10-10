package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "commit the specified files or all outstanding changes",
	Aliases: []string{"ci"},
	GroupID: groups[group_change_creation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().BoolP("addremove", "A", false, "mark new/missing files as added/removed before committing")
	commitCmd.Flags().Bool("amend", false, "amend the parent of the working directory")
	commitCmd.Flags().Bool("close-branch", false, "mark a branch head as closed")
	commitCmd.Flags().StringP("date", "d", "", "record the specified date as commit date")
	commitCmd.Flags().Bool("draft", false, "use the draft phase for committing")
	commitCmd.Flags().BoolP("edit", "e", false, "invoke editor on commit messages")
	commitCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	commitCmd.Flags().Bool("force-close-branch", false, "forcibly close branch from a non-head changeset (ADVANCED)")
	commitCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	commitCmd.Flags().BoolP("interactive", "i", false, "use interactive mode")
	commitCmd.Flags().StringP("logfile", "l", "", "read commit message from file")
	commitCmd.Flags().StringP("message", "m", "", "use text as commit message")
	commitCmd.Flags().BoolP("secret", "s", false, "use the secret phase for committing")
	commitCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	commitCmd.Flags().StringP("user", "u", "", "record the specified user as committer")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"logfile": carapace.ActionFiles(),
	})

	carapace.Gen(commitCmd).PositionalAnyCompletion(
		hg.ActionChangedFiles(),
	)
}
