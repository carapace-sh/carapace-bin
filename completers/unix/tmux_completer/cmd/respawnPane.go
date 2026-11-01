package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var respawnPaneCmd = &cobra.Command{
	Use:     "respawn-pane",
	Aliases: []string{"respawnp"},
	Short:   "reuse a pane in which a command has exited",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(respawnPaneCmd).Standalone()

	respawnPaneCmd.Flags().BoolS("E", "E", false, "leave pane without a running command")
	respawnPaneCmd.Flags().StringS("c", "c", "", "specify a new working directory for the pane")
	respawnPaneCmd.Flags().StringS("e", "e", "", "specify environment variable")
	respawnPaneCmd.Flags().BoolS("k", "k", false, "kill window if it is in use")
	respawnPaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(respawnPaneCmd)

	carapace.Gen(respawnPaneCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(respawnPaneCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
