package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove the specified files on the next commit",
	Aliases: []string{"rm"},
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("after", "A", false, "record delete for missing files")
	removeCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	removeCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	removeCmd.Flags().BoolP("force", "f", false, "forget added files, delete modified files")
	removeCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	removeCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		hg.ActionTrackedFiles(),
	)
}
