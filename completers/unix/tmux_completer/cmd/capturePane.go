package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var capturePaneCmd = &cobra.Command{
	Use:     "capture-pane",
	Aliases: []string{"capturep"},
	Short:   "capture the contents of a pane to a buffer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(capturePaneCmd).Standalone()

	capturePaneCmd.Flags().BoolS("C", "C", false, "escape non-printable characters as octal \\ooo")
	capturePaneCmd.Flags().StringS("E", "E", "", "specify last line to capture")
	capturePaneCmd.Flags().BoolS("F", "F", false, "include flags at the start of each line")
	capturePaneCmd.Flags().BoolS("H", "H", false, "capture only hyperlinks")
	capturePaneCmd.Flags().BoolS("J", "J", false, "join wrapped lines and preserve trailing space")
	capturePaneCmd.Flags().BoolS("L", "L", false, "include line number at the start of each line")
	capturePaneCmd.Flags().BoolS("M", "M", false, "use screen for the mode")
	capturePaneCmd.Flags().BoolS("N", "N", false, "preserve trailing space")
	capturePaneCmd.Flags().BoolS("P", "P", false, "only capture beginnings of as-yet incomplete escape sequences")
	capturePaneCmd.Flags().BoolS("R", "R", false, "dump internal grid data for diagnostics")
	capturePaneCmd.Flags().StringS("S", "S", "", "specify start line to capture")
	capturePaneCmd.Flags().BoolS("a", "a", false, "use alternate screen")
	capturePaneCmd.Flags().StringS("b", "b", "", "choose target buffer")
	capturePaneCmd.Flags().BoolS("e", "e", false, "include escape sequences for attributes etc")
	capturePaneCmd.Flags().BoolS("p", "p", false, "print data to stdout")
	capturePaneCmd.Flags().BoolS("q", "q", false, "ignore errors when trying to access alternate screen")
	capturePaneCmd.Flags().StringS("t", "t", "", "choose source pane")
	rootCmd.AddCommand(capturePaneCmd)

	carapace.Gen(capturePaneCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
		"t": tmux.ActionPanes(),
	})
}
