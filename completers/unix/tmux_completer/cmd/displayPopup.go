package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var displayPopupCmd = &cobra.Command{
	Use:     "display-popup",
	Aliases: []string{"popup"},
	Short:   "display a popup box over a pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayPopupCmd).Standalone()

	displayPopupCmd.Flags().BoolS("B", "B", false, "do not surround the popup by a border")
	displayPopupCmd.Flags().BoolS("C", "C", false, "close any popup on the client")
	displayPopupCmd.Flags().BoolS("E", "E", false, "close the popup when the command exits")
	displayPopupCmd.Flags().BoolS("N", "N", false, "disable any previously specified -E, -EE, or -k")
	displayPopupCmd.Flags().StringS("S", "S", "", "style for the popup border")
	displayPopupCmd.Flags().StringS("T", "T", "", "popup title format")
	displayPopupCmd.Flags().StringS("b", "b", "", "type of border characters")
	displayPopupCmd.Flags().StringS("c", "c", "", "specify target client")
	displayPopupCmd.Flags().StringS("d", "d", "", "specify working directory for the command")
	displayPopupCmd.Flags().StringS("e", "e", "", "set environment variable")
	displayPopupCmd.Flags().StringS("h", "h", "", "specify height")
	displayPopupCmd.Flags().BoolS("k", "k", false, "allow any key to dismiss the popup")
	displayPopupCmd.Flags().StringS("s", "s", "", "style for the popup")
	displayPopupCmd.Flags().StringS("t", "t", "", "specify target pane")
	displayPopupCmd.Flags().StringS("w", "w", "", "specify width")
	displayPopupCmd.Flags().StringS("x", "x", "", "specify horizontal position")
	displayPopupCmd.Flags().StringS("y", "y", "", "specify vertical position")
	rootCmd.AddCommand(displayPopupCmd)

	carapace.Gen(displayPopupCmd).FlagCompletion(carapace.ActionMap{
		"c": tmux.ActionClients(),
		"d": carapace.ActionDirectories(),
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(displayPopupCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
