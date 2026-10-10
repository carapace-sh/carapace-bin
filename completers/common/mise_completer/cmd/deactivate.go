package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Disable mise for current shell session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deactivateCmd).Standalone()

	rootCmd.AddCommand(deactivateCmd)
}
