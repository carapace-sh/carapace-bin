package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var switchClientCmd = &cobra.Command{
	Use:     "switch-client",
	Aliases: []string{"switchc"},
	Short:   "switch the client to another session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchClientCmd).Standalone()

	switchClientCmd.Flags().BoolS("E", "E", false, "don't apply update-environment option")
	switchClientCmd.Flags().StringS("O", "O", "", "initial sort order")
	switchClientCmd.Flags().StringS("T", "T", "", "set the client's key table")
	switchClientCmd.Flags().BoolS("Z", "Z", false, "keep the window zoomed if it was zoomed")
	switchClientCmd.Flags().StringS("c", "c", "", "specify a target client")
	switchClientCmd.Flags().BoolS("l", "l", false, "move client to last session")
	switchClientCmd.Flags().BoolS("n", "n", false, "move client to next session")
	switchClientCmd.Flags().BoolS("p", "p", false, "move client to previous session")
	switchClientCmd.Flags().BoolS("r", "r", false, "toggle read-only flag of client")
	switchClientCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(switchClientCmd)

	carapace.Gen(switchClientCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("name", "size", "creation", "activity"),
		"T": tmux.ActionKeyTables(),
		"c": tmux.ActionClients(),
		"t": tmux.ActionSessions(),
	})
}
