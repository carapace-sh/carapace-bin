package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_help_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install hooks for an AI agent to capture commands in atuin history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_help_installCmd).Standalone()

	hook_helpCmd.AddCommand(hook_help_installCmd)
}
