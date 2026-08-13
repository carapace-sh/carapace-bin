package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Show all domains in a list",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_lsCmd).Standalone()

	domains_lsCmd.Flags().String("format", "", "Output format")
	domains_lsCmd.Flags().Bool("json", false, "Output as JSON")
	domains_lsCmd.Flags().String("limit", "", "Number of results per page")
	domains_lsCmd.Flags().String("next", "", "Show next page of results")

	domainsCmd.AddCommand(domains_lsCmd)

	carapace.Gen(domains_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
