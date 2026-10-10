package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incident_unsubscribeCmd = &cobra.Command{
	Use:     "unsubscribe <id>",
	Short:   "Unsubscribe from an incident.",
	Aliases: []string{"unsub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_unsubscribeCmd).Standalone()

	incidentCmd.AddCommand(incident_unsubscribeCmd)

	carapace.Gen(incident_unsubscribeCmd).PositionalCompletion(
		action.ActionIssues(incident_unsubscribeCmd, "opened"),
	)
}
