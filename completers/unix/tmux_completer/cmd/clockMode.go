package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var clockModeCmd = &cobra.Command{
	Use:   "clock-mode",
	Short: "enter clock mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clockModeCmd).Standalone()

	clockModeCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(clockModeCmd)

	carapace.Gen(clockModeCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
