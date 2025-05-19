package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start a bun (frontend) Dev Server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devCmd).Standalone()

	rootCmd.AddCommand(devCmd)

	carapace.Gen(devCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
