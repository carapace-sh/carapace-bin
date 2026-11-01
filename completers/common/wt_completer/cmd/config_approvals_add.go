package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_approvals_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Store approvals in approvals.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_approvals_addCmd).Standalone()

	config_approvals_addCmd.Flags().Bool("all", false, "Show all commands")
	config_approvals_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_approvalsCmd.AddCommand(config_approvals_addCmd)
}
