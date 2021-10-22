package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Display commands for working with repositories in a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_repositoryCmd).Standalone()
	registryCmd.AddCommand(registry_repositoryCmd)
}
