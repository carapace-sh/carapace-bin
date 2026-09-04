package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Remove all ports of a network from server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_networkCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_networkCmd)
}
