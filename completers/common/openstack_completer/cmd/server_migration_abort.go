package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migration_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Cancel an ongoing live migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_abortCmd).Standalone()

	server_migrationCmd.AddCommand(server_migration_abortCmd)
}
