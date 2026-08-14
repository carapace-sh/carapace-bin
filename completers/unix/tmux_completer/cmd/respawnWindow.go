package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var respawnWindowCmd = &cobra.Command{
	Use:     "respawn-window",
	Aliases: []string{"respawnw"},
	Short:   "reuse a window in which a command has exited",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(respawnWindowCmd).Standalone()

	respawnWindowCmd.Flags().BoolS("E", "E", false, "leave window with one pane and without a running command")
	respawnWindowCmd.Flags().StringS("c", "c", "", "specify a new working directory for the window")
	respawnWindowCmd.Flags().StringS("e", "e", "", "specify environment variable")
	respawnWindowCmd.Flags().BoolS("k", "k", false, "kill window if it is in use")
	respawnWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(respawnWindowCmd)

	carapace.Gen(respawnWindowCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"t": tmux.ActionWindows(),
	})

	carapace.Gen(respawnWindowCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
