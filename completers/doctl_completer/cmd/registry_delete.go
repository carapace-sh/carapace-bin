package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_deleteCmd).Standalone()
	registry_deleteCmd.Flags().BoolP("force", "f", false, "Force registry delete")
	registryCmd.AddCommand(registry_deleteCmd)
}
