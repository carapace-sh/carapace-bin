package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Show all aliases",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_lsCmd).Standalone()

	alias_lsCmd.Flags().String("format", "", "Output format")
	alias_lsCmd.Flags().Bool("json", false, "Output as JSON")
	alias_lsCmd.Flags().String("limit", "", "Number of results per page")
	alias_lsCmd.Flags().String("next", "", "Show next page of results")

	aliasCmd.AddCommand(alias_lsCmd)

	carapace.Gen(alias_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
