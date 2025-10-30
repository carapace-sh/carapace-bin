package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_base_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the workspace (with all applied branches) to include the latest changes from the base branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_base_updateCmd).Standalone()

	help_baseCmd.AddCommand(help_base_updateCmd)
}
