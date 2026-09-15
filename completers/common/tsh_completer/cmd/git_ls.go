package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List Git servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_lsCmd).Standalone()

	git_lsCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	git_lsCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	git_lsCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\").")
	gitCmd.AddCommand(git_lsCmd)
}
