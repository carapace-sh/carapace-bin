package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teams_membersCmd = &cobra.Command{
	Use:     "members",
	Aliases: []string{"member"},
	Short:   "List members for the currently scoped team",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_membersCmd).Standalone()

	teams_membersCmd.Flags().String("format", "", "Output format")
	teams_membersCmd.Flags().Bool("json", false, "Output as JSON")
	teams_membersCmd.Flags().String("limit", "", "Number of results per page")
	teams_membersCmd.Flags().String("next", "", "Show next page of results")

	teamsCmd.AddCommand(teams_membersCmd)
}
