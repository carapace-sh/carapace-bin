package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var killPaneCmd = &cobra.Command{
	Use:     "kill-pane",
	Aliases: []string{"killp"},
	Short:   "destroy a given pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killPaneCmd).Standalone()

	killPaneCmd.Flags().BoolS("a", "a", false, "kill all panes except the one specified by -t")
	killPaneCmd.Flags().StringS("f", "f", "", "filter for panes to kill")
	killPaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(killPaneCmd)

	carapace.Gen(killPaneCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
