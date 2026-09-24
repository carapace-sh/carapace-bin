package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_remove_shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Remove a share from a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_shareCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_shareCmd)
}
