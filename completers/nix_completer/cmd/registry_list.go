package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registry_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available Nix flakes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_listCmd).Standalone()

	registryCmd.AddCommand(registry_listCmd)
}
