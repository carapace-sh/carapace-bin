package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var session_attachCmd = &cobra.Command{
	Use:   "attach <name>",
	Short: "Attach to a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_attachCmd).Standalone()

	sessionCmd.AddCommand(session_attachCmd)

	carapace.Gen(session_attachCmd).PositionalCompletion(action.ActionSessions())
}
