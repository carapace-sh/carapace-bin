package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teams_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Show all teams you're a member of",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_lsCmd).Standalone()

	teams_lsCmd.Flags().String("format", "", "Output format")
	teams_lsCmd.Flags().Bool("json", false, "Output as JSON")
	teams_lsCmd.Flags().String("limit", "", "Number of results per page")
	teams_lsCmd.Flags().String("next", "", "Show next page of results")

	teamsCmd.AddCommand(teams_lsCmd)

	carapace.Gen(teams_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
