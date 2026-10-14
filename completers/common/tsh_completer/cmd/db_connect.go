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

	db_connectCmd.Flags().StringP("db-name", "n", "", "Database name to log in to.")
	db_connectCmd.Flags().StringP("db-roles", "r", "", "List of comma separate database roles to use for auto-provisioned user.")
	db_connectCmd.Flags().StringP("db-user", "u", "", "Database user to log in as.")
	db_connectCmd.Flags().Bool("disable-access-request", false, "Disable automatic resource access requests.")
	db_connectCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	db_connectCmd.Flags().Bool("no-disable-access-request", false, "Disable automatic resource access requests.")
	db_connectCmd.Flags().Bool("no-tunnel", false, "Open authenticated tunnel using database's client certificate so clients don't need to authenticate.")
	db_connectCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	db_connectCmd.Flags().String("request-reason", "", "Reason for requesting access.")
	db_connectCmd.Flags().Bool("tunnel", false, "Open authenticated tunnel using database's client certificate so clients don't need to authenticate.")
	db_connectCmd.Flag("no-disable-access-request").Hidden = true
	db_connectCmd.Flag("no-tunnel").Hidden = true
	db_connectCmd.Flag("tunnel").Hidden = true
	dbCmd.AddCommand(db_connectCmd)
}
