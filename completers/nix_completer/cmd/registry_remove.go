package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove flake from user flake registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_removeCmd).Standalone()

	registry_removeCmd.Flags().String("registry", "", "The registry to operate on")
	registryCmd.AddCommand(registry_removeCmd)

	addLoggingFlags(registry_removeCmd)

	carapace.Gen(registry_removeCmd).FlagCompletion(carapace.ActionMap{
		"registry": carapace.ActionFiles(),
	})

	// TODO positional completion
}
