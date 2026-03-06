package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_approvals_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Store approvals in approvals.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_approvals_addCmd).Standalone()

	hook_approvals_addCmd.Flags().Bool("all", false, "Show all commands")
	hook_approvals_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_approvalsCmd.AddCommand(hook_approvals_addCmd)
}
