package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var showWindowoptionsCmd = &cobra.Command{
	Use:     "show-window-options",
	Aliases: []string{"showw"},
	Short:   "show window options",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showWindowoptionsCmd).Standalone()

	showWindowoptionsCmd.Flags().BoolS("g", "g", false, "show global options")
	showWindowoptionsCmd.Flags().StringS("t", "t", "", "specify target window")
	showWindowoptionsCmd.Flags().BoolS("v", "v", false, "show only the option value, not the name")
	rootCmd.AddCommand(showWindowoptionsCmd)

	carapace.Gen(showWindowoptionsCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
