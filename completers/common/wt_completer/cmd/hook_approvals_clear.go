package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_approvals_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear approved commands from approvals.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_approvals_clearCmd).Standalone()

	hook_approvals_clearCmd.Flags().BoolP("global", "g", false, "Clear global approvals")
	hook_approvals_clearCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_approvalsCmd.AddCommand(hook_approvals_clearCmd)
}
