package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var detachClientCmd = &cobra.Command{
	Use:     "detach-client",
	Aliases: []string{"detach"},
	Short:   "detach a client from the server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detachClientCmd).Standalone()

	detachClientCmd.Flags().StringS("E", "E", "", "run specified shell command to replace the client")
	detachClientCmd.Flags().BoolS("P", "P", false, "send SIGHUP to parent process")
	detachClientCmd.Flags().BoolS("a", "a", false, "kill all clients except for the named by -t")
	detachClientCmd.Flags().StringS("s", "s", "", "specify target session and kill its clients")
	detachClientCmd.Flags().StringS("t", "t", "", "specify target client")
	rootCmd.AddCommand(detachClientCmd)

	carapace.Gen(detachClientCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionSessions(),
		"t": tmux.ActionClients(),
	})
}
