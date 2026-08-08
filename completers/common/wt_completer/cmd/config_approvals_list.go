package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_approvals_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List project commands and their approval status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_approvals_listCmd).Standalone()

	config_approvals_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_approvalsCmd.AddCommand(config_approvals_listCmd)
}
