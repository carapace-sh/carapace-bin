package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove database credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_logoutCmd).Standalone()

	db_logoutCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	db_logoutCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	dbCmd.AddCommand(db_logoutCmd)
}
