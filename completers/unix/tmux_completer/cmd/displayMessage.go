package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var displayMessageCmd = &cobra.Command{
	Use:     "display-message",
	Aliases: []string{"display"},
	Short:   "display a message in the status line",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayMessageCmd).Standalone()

	displayMessageCmd.Flags().BoolS("C", "C", false, "continue updating panes while message is displayed")
	displayMessageCmd.Flags().StringS("F", "F", "", "specify output format")
	displayMessageCmd.Flags().BoolS("I", "I", false, "forward stdin to the target pane")
	displayMessageCmd.Flags().BoolS("N", "N", false, "ignore key presses, close only after delay expires")
	displayMessageCmd.Flags().BoolS("a", "a", false, "list the format variables and their values")
	displayMessageCmd.Flags().StringS("c", "c", "", "specify target client")
	displayMessageCmd.Flags().StringS("d", "d", "", "time to display message")
	displayMessageCmd.Flags().BoolS("l", "l", false, "print message unchanged")
	displayMessageCmd.Flags().BoolS("p", "p", false, "print message to stdout")
	displayMessageCmd.Flags().StringS("t", "t", "", "specify target pane")
	displayMessageCmd.Flags().BoolS("v", "v", false, "print verbose logging as the format is parsed")
	rootCmd.AddCommand(displayMessageCmd)

	carapace.Gen(displayMessageCmd).FlagCompletion(carapace.ActionMap{
		"c": tmux.ActionClients(),
		"t": tmux.ActionPanes(),
	})
}
