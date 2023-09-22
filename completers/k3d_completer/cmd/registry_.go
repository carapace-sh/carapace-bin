package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_Cmd = &cobra.Command{
	Use:   "",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_Cmd).Standalone()

	registryCmd.AddCommand(registry_Cmd)
}
