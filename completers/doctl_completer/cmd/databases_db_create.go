package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_db_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a database within a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_db_createCmd).Standalone()
	databases_dbCmd.AddCommand(databases_db_createCmd)
}
