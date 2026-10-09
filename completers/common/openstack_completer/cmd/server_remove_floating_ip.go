package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_floating_ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Remove floating IP address from server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_floating_ipCmd).Standalone()

	server_remove_floatingCmd.AddCommand(server_remove_floating_ipCmd)
}
