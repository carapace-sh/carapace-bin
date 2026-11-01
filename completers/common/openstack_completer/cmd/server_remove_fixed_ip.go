package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_fixed_ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Remove fixed IP address from server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_fixed_ipCmd).Standalone()

	server_remove_fixedCmd.AddCommand(server_remove_fixed_ipCmd)
}
