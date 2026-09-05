package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var findWindowCmd = &cobra.Command{
	Use:     "find-window",
	Aliases: []string{"findw"},
	Short:   "search for a pattern in windows",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findWindowCmd).Standalone()

	findWindowCmd.Flags().BoolS("C", "C", false, "match only visible window contents")
	findWindowCmd.Flags().BoolS("N", "N", false, "match only the window name")
	findWindowCmd.Flags().BoolS("T", "T", false, "match only the window title")
	findWindowCmd.Flags().BoolS("Z", "Z", false, "zoom the pane")
	findWindowCmd.Flags().BoolS("i", "i", false, "ignore case")
	findWindowCmd.Flags().BoolS("r", "r", false, "use regular expression")
	findWindowCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(findWindowCmd)

	carapace.Gen(findWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
