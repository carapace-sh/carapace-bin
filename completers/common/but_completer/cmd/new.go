package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new TARGET",
	Short:   "Insert a blank commit before the specified commit, or at the top of a stack",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).PositionalCompletion(
		but.ActionLocalBranches(), // TODO commit id
	)
}
