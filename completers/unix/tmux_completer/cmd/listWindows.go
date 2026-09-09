package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var listWindowsCmd = &cobra.Command{
	Use:     "list-windows",
	Aliases: []string{"lsw"},
	Short:   "list windows of a session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listWindowsCmd).Standalone()

	listWindowsCmd.Flags().StringS("F", "F", "", "specify output format")
	listWindowsCmd.Flags().StringS("O", "O", "", "initial sort order")
	listWindowsCmd.Flags().BoolS("a", "a", false, "list all windows the tmux server possesses")
	listWindowsCmd.Flags().StringS("f", "f", "", "filter items")
	listWindowsCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	listWindowsCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(listWindowsCmd)

	carapace.Gen(listWindowsCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("index", "name", "size", "creation", "activity"),
		"t": tmux.ActionSessions(),
	})
}
