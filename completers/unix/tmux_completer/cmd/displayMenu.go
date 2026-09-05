package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var displayMenuCmd = &cobra.Command{
	Use:     "display-menu",
	Aliases: []string{"menu"},
	Short:   "display a menu",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayMenuCmd).Standalone()

	displayMenuCmd.Flags().StringS("C", "C", "", "default selected menu item")
	displayMenuCmd.Flags().StringS("H", "H", "", "style for selected item")
	displayMenuCmd.Flags().BoolS("M", "M", false, "handle mouse events")
	displayMenuCmd.Flags().BoolS("O", "O", false, "don't close menu if mouse is released without making a selection")
	displayMenuCmd.Flags().StringS("S", "S", "", "style for the menu border")
	displayMenuCmd.Flags().StringS("T", "T", "", "specify title")
	displayMenuCmd.Flags().StringS("b", "b", "", "type of border characters")
	displayMenuCmd.Flags().StringS("c", "c", "", "specify target client")
	displayMenuCmd.Flags().StringS("s", "s", "", "style for the menu")
	displayMenuCmd.Flags().StringS("t", "t", "", "specify target pane")
	displayMenuCmd.Flags().StringS("x", "x", "", "specify horizontal position")
	displayMenuCmd.Flags().StringS("y", "y", "", "specify vertical position")
	rootCmd.AddCommand(displayMenuCmd)

	carapace.Gen(displayMenuCmd).FlagCompletion(carapace.ActionMap{
		"c": tmux.ActionClients(),
		"t": tmux.ActionPanes(),
	})
}
