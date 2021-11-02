package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_sqlModeCmd = &cobra.Command{
	Use:   "sql-mode",
	Short: "Display commands to configure a MySQL database cluster's SQL modes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_sqlModeCmd).Standalone()
	databasesCmd.AddCommand(databases_sqlModeCmd)
}
