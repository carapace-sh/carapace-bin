package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nixos-container_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showHostKeyCmd = &cobra.Command{
	Use:   "show-host-key",
	Short: "show the container's SSH host key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showHostKeyCmd).Standalone()

	carapace.Gen(showHostKeyCmd).PositionalCompletion(
		action.ActionContainers(),
	)

	rootCmd.AddCommand(showHostKeyCmd)
}
