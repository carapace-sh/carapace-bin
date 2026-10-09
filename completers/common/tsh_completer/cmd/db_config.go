package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print database connection information. Useful when configuring GUI clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_configCmd).Standalone()

	db_configCmd.Flags().String("db", "", "Print information for the specified database.")
	db_configCmd.Flags().StringP("format", "f", "", "Print format: \"text\" to print in table format (default), \"cmd\" to print connect command, \"json\" or \"yaml\" to print in JSON or YAML.")
	db_configCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	db_configCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	db_configCmd.Flag("db").Hidden = true
	dbCmd.AddCommand(db_configCmd)
}
