package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:     "mark",
	Short:   "Mark a commit or branch for auto-stage or auto-commit",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(markCmd).Standalone()

	markCmd.Flags().BoolP("delete", "d", false, "Deletes a mark")
	markCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(markCmd)

	carapace.Gen(markCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
