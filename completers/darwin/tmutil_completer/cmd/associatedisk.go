package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(associatediskCmd)
}

var associatediskCmd = &cobra.Command{
	Use:   "associatedisk",
	Short: "bind a volume store directory to a local disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(associatediskCmd).Standalone()

	carapace.Gen(associatediskCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}