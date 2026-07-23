package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(session_attachCmd).PositionalCompletion(herdr.ActionSessions())
}
