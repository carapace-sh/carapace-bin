package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Retrieve credentials for a database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_loginCmd).Standalone()

	db_loginCmd.Flags().String("db-name", "", "Optional database name to configure as default.")
	db_loginCmd.Flags().String("db-user", "", "Optional database user to configure as default.")
	dbCmd.AddCommand(db_loginCmd)
}
