package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_approvalsCmd = &cobra.Command{
	Use:    "approvals",
	Short:  "Deprecated: use wt config approvals instead",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_approvalsCmd).Standalone()

	hook_approvalsCmd.Flags().BoolP("help", "h", false, "Print help")
	hookCmd.AddCommand(hook_approvalsCmd)
}
