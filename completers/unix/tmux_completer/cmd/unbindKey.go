package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var unbindKeyCmd = &cobra.Command{
	Use:     "unbind-key",
	Aliases: []string{"unbind"},
	Short:   "unbind a key",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unbindKeyCmd).Standalone()

	unbindKeyCmd.Flags().StringS("T", "T", "", "specify key table")
	unbindKeyCmd.Flags().BoolS("a", "a", false, "remove all key bindings")
	unbindKeyCmd.Flags().BoolS("n", "n", false, "remove a non-prefix binding")
	unbindKeyCmd.Flags().BoolS("q", "q", false, "prevent errors being returned")
	rootCmd.AddCommand(unbindKeyCmd)

	carapace.Gen(unbindKeyCmd).FlagCompletion(carapace.ActionMap{
		"T": tmux.ActionKeyTables(),
	})
}
