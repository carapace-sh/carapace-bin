package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var apps_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_lsCmd).Standalone()

	apps_lsCmd.Flags().BoolP("all", "R", false, "List apps from all clusters and proxies.")
	apps_lsCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	apps_lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	apps_lsCmd.Flags().Bool("no-all", false, "List apps from all clusters and proxies.")
	apps_lsCmd.Flags().Bool("no-verbose", false, "Show extra application fields.")
	apps_lsCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"')")
	apps_lsCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\")")
	apps_lsCmd.Flags().BoolP("verbose", "v", false, "Show extra application fields.")
	apps_lsCmd.Flag("no-all").Hidden = true
	apps_lsCmd.Flag("no-verbose").Hidden = true
	appsCmd.AddCommand(apps_lsCmd)

	// TODO complete cluster
	carapace.Gen(apps_lsCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"format":  tsh.ActionFormats(),
	})
}
