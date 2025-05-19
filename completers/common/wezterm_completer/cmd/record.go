package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var recordCmd = &cobra.Command{
	Use:   "record [PROG]...",
	Short: "Record a terminal session as an asciicast",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordCmd).Standalone()
	recordCmd.Flags().SetInterspersed(false)

	recordCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(recordCmd)

	carapace.Gen(recordCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
