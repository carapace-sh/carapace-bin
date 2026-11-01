package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nixos-container_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	carapace.Gen(stopCmd).PositionalCompletion(
		action.ActionContainers(),
	)

	rootCmd.AddCommand(stopCmd)
}
