package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var sendKeysCmd = &cobra.Command{
	Use:     "send-keys",
	Aliases: []string{"send"},
	Short:   "send key(s) to a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendKeysCmd).Standalone()

	sendKeysCmd.Flags().BoolS("F", "F", false, "expand formats in arguments where appropriate")
	sendKeysCmd.Flags().BoolS("H", "H", false, "interpret key as hexadecimal number for an ASCII character")
	sendKeysCmd.Flags().BoolS("K", "K", false, "send keys to target-client for key table lookup")
	sendKeysCmd.Flags().BoolS("M", "M", false, "pass through a mouse event")
	sendKeysCmd.Flags().StringS("N", "N", "", "specify repeat count")
	sendKeysCmd.Flags().BoolS("R", "R", false, "reset terminal state")
	sendKeysCmd.Flags().BoolS("X", "X", false, "send a command into copy mode")
	sendKeysCmd.Flags().StringS("c", "c", "", "specify target client")
	sendKeysCmd.Flags().BoolS("l", "l", false, "disable key name lookup and send data literally")
	sendKeysCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(sendKeysCmd)

	carapace.Gen(sendKeysCmd).FlagCompletion(carapace.ActionMap{
		"c": tmux.ActionClients(),
		"t": tmux.ActionPanes(),
	})
}
