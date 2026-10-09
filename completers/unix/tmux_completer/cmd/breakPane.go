package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var breakPaneCmd = &cobra.Command{
	Use:     "break-pane",
	Aliases: []string{"breakp"},
	Short:   "break a pane from an existing into a new window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(breakPaneCmd).Standalone()

	breakPaneCmd.Flags().StringS("F", "F", "", "specify output format")
	breakPaneCmd.Flags().BoolS("P", "P", false, "print information of new window after it has been created")
	breakPaneCmd.Flags().BoolS("W", "W", false, "make src-pane floating")
	breakPaneCmd.Flags().StringS("X", "X", "", "x position of floating pane")
	breakPaneCmd.Flags().StringS("Y", "Y", "", "y position of floating pane")
	breakPaneCmd.Flags().BoolS("a", "a", false, "move window to next index after")
	breakPaneCmd.Flags().BoolS("b", "b", false, "move window to next index before")
	breakPaneCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	breakPaneCmd.Flags().StringS("n", "n", "", "specify window name")
	breakPaneCmd.Flags().StringS("s", "s", "", "specify source pane")
	breakPaneCmd.Flags().StringS("t", "t", "", "specify destination window")
	breakPaneCmd.Flags().StringS("x", "x", "", "width of floating pane")
	breakPaneCmd.Flags().StringS("y", "y", "", "height of floating pane")
	rootCmd.AddCommand(breakPaneCmd)

	carapace.Gen(breakPaneCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionPanes(),
		"t": tmux.ActionWindows(),
	})
}
