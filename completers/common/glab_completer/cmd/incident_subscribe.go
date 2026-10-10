package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incident_subscribeCmd = &cobra.Command{
	Use:     "subscribe <id>",
	Short:   "Subscribe to an incident.",
	Aliases: []string{"sub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_subscribeCmd).Standalone()

	incidentCmd.AddCommand(incident_subscribeCmd)

	carapace.Gen(incident_subscribeCmd).PositionalCompletion(
		action.ActionIssues(incident_subscribeCmd, "opened"),
	)
}
