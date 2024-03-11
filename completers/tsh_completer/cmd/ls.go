package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List remote SSH nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().BoolP("all", "R", false, "List nodes from all clusters and proxies.")
	lsCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml, names)")
	lsCmd.Flags().Bool("no-all", false, "List nodes from all clusters and proxies.")
	lsCmd.Flags().Bool("no-verbose", false, "One-line output (for text format), including node UUIDs")
	lsCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"')")
	lsCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\")")
	lsCmd.Flags().BoolP("verbose", "v", false, "One-line output (for text format), including node UUIDs")
	lsCmd.Flag("no-all").Hidden = true
	lsCmd.Flag("no-verbose").Hidden = true
	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"format":  tsh.ActionFormats(),
	})
}
