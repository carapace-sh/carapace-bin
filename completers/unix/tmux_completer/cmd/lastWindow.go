package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var lastWindowCmd = &cobra.Command{
	Use:     "last-window",
	Aliases: []string{"last"},
	Short:   "select the previously selected window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lastWindowCmd).Standalone()

	lastWindowCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(lastWindowCmd)

	carapace.Gen(lastWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
