package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_trackCmd = &cobra.Command{
	Use:   "track [OPTIONS] <PATHS>...",
	Short: "Start tracking specified paths in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_trackCmd).Standalone()

	file_trackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_trackCmd.Flags().Bool("include-ignored", false, "Track paths even if they're ignored or too large")
	fileCmd.AddCommand(file_trackCmd)

	carapace.Gen(file_trackCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
