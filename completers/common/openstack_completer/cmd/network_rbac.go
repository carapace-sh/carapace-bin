package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_rbacCmd = &cobra.Command{
	Use:   "rbac",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rbacCmd).Standalone()

	networkCmd.AddCommand(network_rbacCmd)
}
