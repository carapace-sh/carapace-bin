package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerRegistry_repositoryCmd = &cobra.Command{
	Use:   "repository <command> [flags]",
	Short: "Manage container registry repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_repositoryCmd).Standalone()

	containerRegistryCmd.AddCommand(containerRegistry_repositoryCmd)
}
