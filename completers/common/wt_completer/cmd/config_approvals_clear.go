package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_approvals_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear approved commands from approvals.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_approvals_clearCmd).Standalone()

	config_approvals_clearCmd.Flags().BoolP("global", "g", false, "Clear global approvals")
	config_approvals_clearCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_approvals_clearCmd.Flags().Bool("stale", false, "Clear only stale approvals")
	config_approvalsCmd.AddCommand(config_approvals_clearCmd)
}
