package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incident_closeCmd = &cobra.Command{
	Use:     "close [<id> | <url>] [flags]",
	Short:   "Close an incident.",
	Aliases: []string{"resolve"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_closeCmd).Standalone()

	incidentCmd.AddCommand(incident_closeCmd)

	carapace.Gen(incident_closeCmd).PositionalCompletion(
		action.ActionIssues(incident_closeCmd, "opened"),
	)
}
