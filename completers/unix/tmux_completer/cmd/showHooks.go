package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var showHooksCmd = &cobra.Command{
	Use:   "show-hooks",
	Short: "show the global list of hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showHooksCmd).Standalone()

	showHooksCmd.Flags().BoolS("B", "B", false, "show subscriptions installed with set-hook -B")
	showHooksCmd.Flags().StringS("F", "F", "", "specify format for each hook")
	showHooksCmd.Flags().BoolS("g", "g", false, "show global hooks")
	showHooksCmd.Flags().BoolS("p", "p", false, "show pane hooks")
	showHooksCmd.Flags().StringS("t", "t", "", "specify target pane")
	showHooksCmd.Flags().BoolS("w", "w", false, "show window hooks")
	rootCmd.AddCommand(showHooksCmd)

	carapace.Gen(showHooksCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
