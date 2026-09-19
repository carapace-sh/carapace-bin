package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var forgetCmd = &cobra.Command{
	Use:     "forget",
	Short:   "forget the specified files on the next commit",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forgetCmd).Standalone()

	forgetCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	forgetCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	forgetCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	forgetCmd.Flags().BoolP("interactive", "i", false, "use interactive mode")
	rootCmd.AddCommand(forgetCmd)

	carapace.Gen(forgetCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})

	carapace.Gen(forgetCmd).PositionalAnyCompletion(
		hg.ActionTrackedFiles(),
	)
}
