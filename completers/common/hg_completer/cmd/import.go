package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Short:   "import an ordered set of patches",
	Aliases: []string{"patch"},
	GroupID: groups[group_change_import_export].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().StringP("base", "b", "", "base path (DEPRECATED)")
	importCmd.Flags().Bool("bypass", false, "apply patch without touching the working directory")
	importCmd.Flags().StringP("date", "d", "", "record the specified date as commit date")
	importCmd.Flags().BoolP("edit", "e", false, "invoke editor on commit messages")
	importCmd.Flags().Bool("exact", false, "abort if patch would apply lossily")
	importCmd.Flags().BoolP("force", "f", false, "skip check for outstanding uncommitted changes (DEPRECATED)")
	importCmd.Flags().Bool("import-branch", false, "use any branch information in patch (implied by --exact)")
	importCmd.Flags().StringP("logfile", "l", "", "read commit message from file")
	importCmd.Flags().StringP("message", "m", "", "use text as commit message")
	importCmd.Flags().Bool("no-commit", false, "don't commit, just update the working directory")
	importCmd.Flags().Bool("partial", false, "commit even if some hunks fail")
	importCmd.Flags().String("prefix", "", "apply patch to subdirectory")
	importCmd.Flags().Bool("secret", false, "use the secret phase for committing")
	importCmd.Flags().StringP("similarity", "s", "", "guess renamed files by similarity (0<=s<=100)")
	importCmd.Flags().StringP("strip", "p", "", "directory strip option for patch.")
	importCmd.Flags().StringP("user", "u", "", "record the specified user as committer")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"base":    action.ActionRevisions(),
		"logfile": carapace.ActionFiles(),
	})

	carapace.Gen(importCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
