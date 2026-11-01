package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nixos-container_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "execute a command inside the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	carapace.Gen(runCmd).PositionalCompletion(
		action.ActionContainers(),
	)

	rootCmd.AddCommand(runCmd)
}
