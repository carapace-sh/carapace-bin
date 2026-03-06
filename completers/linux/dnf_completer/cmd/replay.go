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

	replayCmd.Flags().Bool("ignore-extras", false, "don't consider extra packages as errors")
	replayCmd.Flags().Bool("ignore-installed", false, "don't consider mismatches as errors")
	replayCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	replayCmd.Flags().Bool("skip-unavailable", false, "skip unavailable packages")

	rootCmd.AddCommand(replayCmd)

	carapace.Gen(replayCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
