package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incident_noteCmd = &cobra.Command{
	Use:     "note <incident-id>",
	Short:   "Comment on an incident in GitLab.",
	Aliases: []string{"comment"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_noteCmd).Standalone()

	incident_noteCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the comment. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	incident_noteCmd.Flags().StringP("message", "m", "", "Message text.")
	incidentCmd.AddCommand(incident_noteCmd)

	carapace.Gen(incident_noteCmd).PositionalCompletion(
		action.ActionIssues(incident_noteCmd, "opened"),
	)
}
