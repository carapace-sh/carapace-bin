package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_connectCmd).Standalone()

	db_connectCmd.Flags().String("db-name", "", "Optional database name to log in to.")
	db_connectCmd.Flags().String("db-user", "", "Optional database user to log in as.")
	dbCmd.AddCommand(db_connectCmd)
}
