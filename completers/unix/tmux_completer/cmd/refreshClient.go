package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var refreshClientCmd = &cobra.Command{
	Use:     "refresh-client",
	Aliases: []string{"refresh"},
	Short:   "refresh a client",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshClientCmd).Standalone()

	refreshClientCmd.Flags().StringS("A", "A", "", "allow a control mode client to trigger actions on a pane")
	refreshClientCmd.Flags().StringS("B", "B", "", "set a subscription to a format for a control mode client")
	refreshClientCmd.Flags().StringS("C", "C", "", "set the width and height of a control client")
	refreshClientCmd.Flags().BoolS("D", "D", false, "move visible portion of window down")
	refreshClientCmd.Flags().BoolS("L", "L", false, "move visible portion of window left")
	refreshClientCmd.Flags().BoolS("R", "R", false, "move visible portion of window right")
	refreshClientCmd.Flags().BoolS("S", "S", false, "only update the client's status bar")
	refreshClientCmd.Flags().BoolS("U", "U", false, "move visible portion of window up")
	refreshClientCmd.Flags().BoolS("c", "c", false, "reset so that the position follows the cursor")
	refreshClientCmd.Flags().StringS("f", "f", "", "set client flags")
	refreshClientCmd.Flags().BoolS("l", "l", false, "request clipboard from the client")
	refreshClientCmd.Flags().StringS("r", "r", "", "provide information about a pane via a report")
	refreshClientCmd.Flags().StringS("t", "t", "", "specify target client")
	rootCmd.AddCommand(refreshClientCmd)

	carapace.Gen(refreshClientCmd).FlagCompletion(carapace.ActionMap{
		"f": tmux.ActionClientFlags().UniqueList(","),
		"t": tmux.ActionClients(),
	})
}
