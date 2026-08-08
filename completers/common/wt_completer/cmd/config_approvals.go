package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_approvalsCmd = &cobra.Command{
	Use:   "approvals",
	Short: "Manage command approvals",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_approvalsCmd).Standalone()

	config_approvalsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_approvalsCmd)
}
