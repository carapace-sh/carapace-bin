package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rebuildCmd).Standalone()

	helpCmd.AddCommand(help_rebuildCmd)
}
