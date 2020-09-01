package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).PositionalAnyCompletion(ActionServices())
}
