package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forwardNewCmd = &cobra.Command{
	Use:   "new (-L|-R) [bind_address:]port:host:hostport...",
	Short: "Forward ports until interrupted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardNewCmd).Standalone()

	forwardNewCmd.Flags().StringArrayP("local", "L", nil, "listen here on port and connect to host:hostport from the server")
	forwardNewCmd.Flags().StringArrayP("remote", "R", nil, "listen on the server on port and connect to host:hostport from here")
	forwardNewCmd.Flags().StringP("session", "s", "", "keep the forwards' blocks in session SESSION")

	forwardCmd.AddCommand(forwardNewCmd)

	carapace.Gen(forwardNewCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(forwardNewCmd),
	})
}
