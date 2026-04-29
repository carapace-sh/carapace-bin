package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hook_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install hooks for an AI agent to capture commands in atuin history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hook_installCmd).Standalone()

	help_hookCmd.AddCommand(help_hook_installCmd)
}
