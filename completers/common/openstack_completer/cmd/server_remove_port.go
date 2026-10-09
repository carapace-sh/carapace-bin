package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_portCmd = &cobra.Command{
	Use:   "port",
	Short: "Remove port from server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_portCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_portCmd)
}
