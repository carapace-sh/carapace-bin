package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var providesCmd = &cobra.Command{
	Use:   "provides",
	Short: "Determine which port owns a given file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providesCmd).Standalone()
	rootCmd.AddCommand(providesCmd)
	carapace.Gen(providesCmd).PositionalCompletion(carapace.ActionFiles())
}
