package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:     "revert",
	Short:   "restore files to their checkout state",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revertCmd).Standalone()

	revertCmd.Flags().BoolP("all", "a", false, "revert all changes when no arguments given")
	revertCmd.Flags().StringP("date", "d", "", "tipmost revision matching date")
	revertCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	revertCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	revertCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	revertCmd.Flags().BoolP("interactive", "i", false, "interactively select the changes")
	revertCmd.Flags().BoolP("no-backup", "C", false, "do not save backup copies of files")
	revertCmd.Flags().StringP("rev", "r", "", "revert to the specified revision")
	rootCmd.AddCommand(revertCmd)

	carapace.Gen(revertCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     action.ActionRevisions(),
	})

	carapace.Gen(revertCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
