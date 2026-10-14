package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migration_confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Confirm server migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_confirmCmd).Standalone()

	server_migrationCmd.AddCommand(server_migration_confirmCmd)
}
