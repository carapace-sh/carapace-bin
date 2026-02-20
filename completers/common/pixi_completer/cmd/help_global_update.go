package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates environments in the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_updateCmd).Standalone()

	help_globalCmd.AddCommand(help_global_updateCmd)
}
