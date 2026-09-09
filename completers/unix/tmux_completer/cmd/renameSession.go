package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var renameSessionCmd = &cobra.Command{
	Use:     "rename-session",
	Aliases: []string{"rename"},
	Short:   "rename a session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameSessionCmd).Standalone()

	renameSessionCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(renameSessionCmd)

	carapace.Gen(renameSessionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
