package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var showOptionsCmd = &cobra.Command{
	Use:     "show-options",
	Aliases: []string{"show"},
	Short:   "show session options",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showOptionsCmd).Standalone()

	showOptionsCmd.Flags().BoolS("A", "A", false, "include options inherited from a parent set")
	showOptionsCmd.Flags().StringS("F", "F", "", "specify format")
	showOptionsCmd.Flags().BoolS("H", "H", false, "include hooks")
	showOptionsCmd.Flags().BoolS("g", "g", false, "show global options")
	showOptionsCmd.Flags().BoolS("p", "p", false, "show pane (not session) options")
	showOptionsCmd.Flags().BoolS("q", "q", false, "suppress errors about unknown or ambiguous options")
	showOptionsCmd.Flags().BoolS("s", "s", false, "show server (not session) options")
	showOptionsCmd.Flags().StringS("t", "t", "", "specify target session")
	showOptionsCmd.Flags().BoolS("v", "v", false, "show only the option value, not the name")
	showOptionsCmd.Flags().BoolS("w", "w", false, "show window (not session) options")
	rootCmd.AddCommand(showOptionsCmd)

	carapace.Gen(showOptionsCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
