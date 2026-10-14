package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var linkWindowCmd = &cobra.Command{
	Use:     "link-window",
	Aliases: []string{"linkw"},
	Short:   "link a window to another",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkWindowCmd).Standalone()

	linkWindowCmd.Flags().BoolS("a", "a", false, "move window to next index after destination window")
	linkWindowCmd.Flags().BoolS("b", "b", false, "move window to next index before destination window")
	linkWindowCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	linkWindowCmd.Flags().BoolS("k", "k", false, "kill the target window if it exists")
	linkWindowCmd.Flags().StringS("s", "s", "", "specify source window")
	linkWindowCmd.Flags().StringS("t", "t", "", "specify destination window")
	rootCmd.AddCommand(linkWindowCmd)

	carapace.Gen(linkWindowCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionWindows(),
		"t": tmux.ActionWindows(),
	})
}
