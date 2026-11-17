package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forwardOpenCmd = &cobra.Command{
	Use:   "open [flags] (URL | [host:]port | BLOCK)",
	Short: "Forward a local port to a server-side target and open it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardOpenCmd).Standalone()

	forwardOpenCmd.Flags().StringP("bind", "b", "127.0.0.1", "listen on ADDRESS; * means every interface")
	forwardOpenCmd.Flags().Bool("no-open", false, "print the URL without opening it")
	forwardOpenCmd.Flags().StringP("port", "p", "", "listen on PORT; 0 picks a free port")
	forwardOpenCmd.Flags().StringP("session", "s", "", "keep the forward's block in session SESSION")

	forwardCmd.AddCommand(forwardOpenCmd)

	carapace.Gen(forwardOpenCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(forwardOpenCmd),
	})

	carapace.Gen(forwardOpenCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionBlocks(forwardOpenCmd),
		).ToA(),
	)
}
