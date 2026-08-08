package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var replayCmd = &cobra.Command{
	Use:   "replay [options] <transaction-path>",
	Short: "replay stored transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replayCmd).Standalone()

	replayCmd.Flags().Bool("ignore-extras", false, "Ignore extra packages")
	replayCmd.Flags().Bool("ignore-installed", false, "Ignore installed packages")
	replayCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	replayCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")

	rootCmd.AddCommand(replayCmd)

	carapace.Gen(replayCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
