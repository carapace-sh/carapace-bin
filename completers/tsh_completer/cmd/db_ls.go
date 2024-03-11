package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var db_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available databases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_lsCmd).Standalone()

	db_lsCmd.Flags().BoolP("all", "R", false, "List databases from all clusters and proxies.")
	db_lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	db_lsCmd.Flags().Bool("no-all", false, "List databases from all clusters and proxies.")
	db_lsCmd.Flags().Bool("no-verbose", false, "Show extra database fields.")
	db_lsCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"')")
	db_lsCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\")")
	db_lsCmd.Flags().BoolP("verbose", "v", false, "Show extra database fields.")
	db_lsCmd.Flag("no-all").Hidden = true
	db_lsCmd.Flag("no-verbose").Hidden = true
	dbCmd.AddCommand(db_lsCmd)

	carapace.Gen(db_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
