package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_base_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates all applied branches to be up to date with the target branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_base_updateCmd).Standalone()

	help_baseCmd.AddCommand(help_base_updateCmd)
}
