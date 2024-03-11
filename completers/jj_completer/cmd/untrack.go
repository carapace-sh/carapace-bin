package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untrackCmd = &cobra.Command{
	Use:   "untrack [OPTIONS] <PATHS>...",
	Short: "Stop tracking specified paths in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untrackCmd).Standalone()

	untrackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(untrackCmd)

	carapace.Gen(untrackCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
