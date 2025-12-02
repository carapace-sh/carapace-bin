package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var rubCmd = &cobra.Command{
	Use:     "rub",
	Short:   "Combines two entities together to perform an operation like amend, squash, assign, or move",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(rubCmd).Standalone()

	rubCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(rubCmd)

	carapace.Gen(rubCmd).PositionalCompletion(
		but.ActionLocalBranches(),
		but.ActionLocalBranches().FilterArgs(),
	)
}
