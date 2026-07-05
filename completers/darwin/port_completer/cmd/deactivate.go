package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate the installed port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deactivateCmd).Standalone()
	rootCmd.AddCommand(deactivateCmd)
}
