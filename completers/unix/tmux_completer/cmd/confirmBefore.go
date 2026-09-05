package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var confirmBeforeCmd = &cobra.Command{
	Use:     "confirm-before",
	Aliases: []string{"confirm"},
	Short:   "run a command but ask for confirmation before",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(confirmBeforeCmd).Standalone()

	confirmBeforeCmd.Flags().BoolS("b", "b", false, "show prompt in background")
	confirmBeforeCmd.Flags().StringS("c", "c", "", "change the confirmation key")
	confirmBeforeCmd.Flags().StringS("p", "p", "", "specify prompt")
	confirmBeforeCmd.Flags().StringS("t", "t", "", "specify target client")
	confirmBeforeCmd.Flags().BoolS("y", "y", false, "run command on Enter")
	rootCmd.AddCommand(confirmBeforeCmd)

	carapace.Gen(confirmBeforeCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionClients(),
	})

	carapace.Gen(confirmBeforeCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("tmux").Shift(1),
	)
}
