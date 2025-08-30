package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "generates a completion for the provided input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completeCmd).Standalone()

	completeCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.AddCommand(completeCmd)

	carapace.Gen(completeCmd).PositionalCompletion(
		bridge.ActionCarapaceBin().Split(),
	)
}
