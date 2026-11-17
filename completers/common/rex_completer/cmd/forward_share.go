package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forwardShareCmd = &cobra.Command{
	Use:   "share [flags] (URL | [host:]port)",
	Short: "Forward a server port to a target on this machine's side",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardShareCmd).Standalone()

	forwardShareCmd.Flags().StringP("bind", "b", "127.0.0.1", "listen on ADDRESS; * means every interface")
	forwardShareCmd.Flags().StringP("port", "p", "", "listen on PORT; 0 picks a free port")
	forwardShareCmd.Flags().StringP("session", "s", "", "keep the forward's block in session SESSION")

	forwardCmd.AddCommand(forwardShareCmd)

	carapace.Gen(forwardShareCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(forwardShareCmd),
	})
}
