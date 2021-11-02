package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Display commands for managing individual databases within a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_dbCmd).Standalone()
	databasesCmd.AddCommand(databases_dbCmd)
}
