package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "add the specified files on the next commit",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	addCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	addCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	addCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.Batch(action.ActionUntrackedFiles(), carapace.ActionFiles()).ToA(),
	)
}
