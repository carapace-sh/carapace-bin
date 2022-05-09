package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var killSessionCmd = &cobra.Command{
	Use:   "kill-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killSessionCmd).Standalone()

	killSessionCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(killSessionCmd)

	carapace.Gen(killSessionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
