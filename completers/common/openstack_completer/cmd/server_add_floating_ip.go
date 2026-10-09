package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_floating_ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Add floating IP address to server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_floating_ipCmd).Standalone()

	server_add_floating_ipCmd.Flags().String("fixed-ip-address", "", "Fixed IP address to associate with this floating IP address.")
	server_add_floatingCmd.AddCommand(server_add_floating_ipCmd)
}
