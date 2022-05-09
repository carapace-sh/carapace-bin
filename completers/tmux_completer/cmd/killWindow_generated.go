package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var killWindowCmd = &cobra.Command{
	Use:   "kill-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killWindowCmd).Standalone()

	killWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	killWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(killWindowCmd)

	carapace.Gen(killWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
