package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install hooks for an AI agent to capture commands in atuin history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_installCmd).Standalone()

	hook_installCmd.Flags().BoolP("help", "h", false, "Print help")
	hookCmd.AddCommand(hook_installCmd)

	carapace.Gen(hook_installCmd).PositionalCompletion(
		carapace.ActionValues("claude-code", "codex", "pi"),
	)
}
