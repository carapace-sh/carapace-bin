package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_restoreCmd).Standalone()

	serverCmd.AddCommand(server_restoreCmd)
}
