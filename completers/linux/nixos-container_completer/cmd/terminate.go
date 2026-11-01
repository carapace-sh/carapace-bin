package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nixos-container_completer/cmd/action"
	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "forcefully terminate a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminateCmd).Standalone()

	carapace.Gen(terminateCmd).PositionalCompletion(
		action.ActionContainers(),
	)

	rootCmd.AddCommand(terminateCmd)
}
