package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nixos-container_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showIpCmd = &cobra.Command{
	Use:   "show-ip",
	Short: "show the container's IP address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showIpCmd).Standalone()

	carapace.Gen(showIpCmd).PositionalCompletion(
		action.ActionContainers(),
	)

	rootCmd.AddCommand(showIpCmd)
}
