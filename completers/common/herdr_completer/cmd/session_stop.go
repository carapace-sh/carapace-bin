package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var session_stopCmd = &cobra.Command{
	Use:   "stop <name>",
	Short: "Stop a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_stopCmd).Standalone()

	session_stopCmd.Flags().Bool("json", false, "")
	sessionCmd.AddCommand(session_stopCmd)

	carapace.Gen(session_stopCmd).PositionalCompletion(action.ActionSessions())
}
