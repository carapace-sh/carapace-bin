package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nixos-container_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show container status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	carapace.Gen(statusCmd).PositionalCompletion(
		action.ActionContainers(),
	)

	rootCmd.AddCommand(statusCmd)
}
