package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migrate_confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "DEPRECATED: Use 'server migration confirm' instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migrate_confirmCmd).Standalone()

	server_migrateCmd.AddCommand(server_migrate_confirmCmd)
}
