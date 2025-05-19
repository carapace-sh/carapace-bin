package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var teams_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_switchCmd).Standalone()

	teamsCmd.AddCommand(teams_switchCmd)

	carapace.Gen(teams_switchCmd).PositionalCompletion(
		action.ActionTeams(teams_switchCmd),
	)
}
