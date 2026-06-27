package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/crush"
	"github.com/spf13/cobra"
)

var session_showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Show session details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_showCmd).Standalone()

	session_showCmd.Flags().Bool("json", false, "output in JSON format")
	sessionCmd.AddCommand(session_showCmd)

	carapace.Gen(session_showCmd).PositionalCompletion(
		crush.ActionSessions(),
	)
}
