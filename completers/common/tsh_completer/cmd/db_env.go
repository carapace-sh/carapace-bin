package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print environment variables for the configured database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_envCmd).Standalone()

	db_envCmd.Flags().String("db", "", "Print environment for the specified database.")
	db_envCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	db_envCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	db_envCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	db_envCmd.Flag("db").Hidden = true
	dbCmd.AddCommand(db_envCmd)
}
