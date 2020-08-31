package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(pauseCmd)

	carapace.Gen(pauseCmd).PositionalAnyCompletion(ActionServices())
}
