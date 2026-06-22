package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbMigrateCmd).Standalone()
	dbMigrateCmd.Flags().BoolP("help", "h", false, "Show help")
	dbMigrateCmd.Flags().Int("version", 0, "Migrate to version (VERSION=)")
}
