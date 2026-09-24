package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var unshelveCmd = &cobra.Command{
	Use:     "unshelve",
	Short:   "restore a shelved change to the working directory",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unshelveCmd).Standalone()

	unshelveCmd.Flags().BoolP("abort", "a", false, "abort an incomplete unshelve operation")
	unshelveCmd.Flags().BoolP("continue", "c", false, "continue an incomplete unshelve operation")
	unshelveCmd.Flags().String("date", "", "set date for temporary commits (DEPRECATED)")
	unshelveCmd.Flags().BoolP("interactive", "i", false, "use interactive mode (EXPERIMENTAL)")
	unshelveCmd.Flags().BoolP("keep", "k", false, "keep shelve after unshelving")
	unshelveCmd.Flags().StringP("name", "n", "", "restore shelved change with given name")
	unshelveCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	rootCmd.AddCommand(unshelveCmd)

	carapace.Gen(unshelveCmd).FlagCompletion(carapace.ActionMap{
		"name": hg.ActionShelves(),
		"tool": hg.ActionMergeTools(),
	})

	carapace.Gen(unshelveCmd).PositionalCompletion(
		hg.ActionShelves(),
	)
}
