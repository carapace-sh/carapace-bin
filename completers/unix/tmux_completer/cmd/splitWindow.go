package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var splitWindowCmd = &cobra.Command{
	Use:     "split-window",
	Aliases: []string{"splitw"},
	Short:   "splits a pane into two",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitWindowCmd).Standalone()

	splitWindowCmd.Flags().StringS("B", "B", "", "set pane border lines")
	splitWindowCmd.Flags().StringS("F", "F", "", "specify output format")
	splitWindowCmd.Flags().BoolS("I", "I", false, "create empty pane and forward stdin to it")
	splitWindowCmd.Flags().BoolS("P", "P", false, "print information of new pane after it has been created")
	splitWindowCmd.Flags().StringS("R", "R", "", "border style when inactive")
	splitWindowCmd.Flags().StringS("S", "S", "", "border style when active")
	splitWindowCmd.Flags().StringS("T", "T", "", "pane title")
	splitWindowCmd.Flags().BoolS("Z", "Z", false, "zoom the pane")
	splitWindowCmd.Flags().BoolS("b", "b", false, "create new pane left of or above target pane")
	splitWindowCmd.Flags().StringS("c", "c", "", "specify working directory for the session")
	splitWindowCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	splitWindowCmd.Flags().StringS("e", "e", "", "specify environment variable")
	splitWindowCmd.Flags().BoolS("f", "f", false, "create new pane spanning full window width or height")
	splitWindowCmd.Flags().BoolS("h", "h", false, "split horizontally")
	splitWindowCmd.Flags().BoolS("k", "k", false, "keep pane open after command exits")
	splitWindowCmd.Flags().StringS("l", "l", "", "define new pane's size")
	splitWindowCmd.Flags().StringS("m", "m", "", "equivalent to -k but also sets remain-on-exit-format")
	splitWindowCmd.Flags().StringS("p", "p", "", "shorthand for size as percentage")
	splitWindowCmd.Flags().StringS("s", "s", "", "style for pane content")
	splitWindowCmd.Flags().StringS("t", "t", "", "specify target pane")
	splitWindowCmd.Flags().BoolS("v", "v", false, "split vertically")
	rootCmd.AddCommand(splitWindowCmd)

	carapace.Gen(splitWindowCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(splitWindowCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
