package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_migrationCmd).Standalone()

	share_serverCmd.AddCommand(share_server_migrationCmd)
}
