package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_untrackCmd = &cobra.Command{
	Use:   "untrack [OPTIONS] <PATHS>...",
	Short: "Stop tracking specified paths in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_untrackCmd).Standalone()

	file_untrackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	fileCmd.AddCommand(file_untrackCmd)

	carapace.Gen(file_untrackCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
