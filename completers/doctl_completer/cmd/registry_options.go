package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "List available container registry options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_optionsCmd).Standalone()
	registryCmd.AddCommand(registry_optionsCmd)
}
