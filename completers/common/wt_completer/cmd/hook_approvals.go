package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_approvalsCmd = &cobra.Command{
	Use:   "approvals",
	Short: "Manage command approvals",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_approvalsCmd).Standalone()

	hook_approvalsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hookCmd.AddCommand(hook_approvalsCmd)
}
