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

	db_loginCmd.Flags().StringP("db-name", "n", "", "Database name to configure as default.")
	db_loginCmd.Flags().StringP("db-roles", "r", "", "List of comma separate database roles to use for auto-provisioned user.")
	db_loginCmd.Flags().StringP("db-user", "u", "", "Database user to configure as default.")
	db_loginCmd.Flags().Bool("disable-access-request", false, "Disable automatic resource access requests.")
	db_loginCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	db_loginCmd.Flags().Bool("no-disable-access-request", false, "Disable automatic resource access requests.")
	db_loginCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	db_loginCmd.Flags().String("request-reason", "", "Reason for requesting access.")
	db_loginCmd.Flag("no-disable-access-request").Hidden = true
	dbCmd.AddCommand(db_loginCmd)
}
