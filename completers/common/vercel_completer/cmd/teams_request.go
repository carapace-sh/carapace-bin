package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teams_requestCmd = &cobra.Command{
	Use:     "request",
	Aliases: []string{"access-request"},
	Short:   "Show join-request status for the current team",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_requestCmd).Standalone()

	teams_requestCmd.Flags().String("format", "", "Output format")
	teams_requestCmd.Flags().Bool("json", false, "Output as JSON")

	teamsCmd.AddCommand(teams_requestCmd)

	carapace.Gen(teams_requestCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
