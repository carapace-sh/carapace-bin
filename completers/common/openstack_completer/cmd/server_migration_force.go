package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migration_forceCmd = &cobra.Command{
	Use:   "force",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_forceCmd).Standalone()

	server_migrationCmd.AddCommand(server_migration_forceCmd)
}
