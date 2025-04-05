package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time one of the other commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timeCmd).Standalone()
	timeCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(timeCmd)

	carapace.Gen(timeCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gm"),
	)
}
