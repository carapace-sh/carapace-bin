package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var powerSaverCmd = &cobra.Command{
	Use:   "power-saver",
	Short: "Apply the power-saver profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(powerSaverCmd).Standalone()

	rootCmd.AddCommand(powerSaverCmd)
}
