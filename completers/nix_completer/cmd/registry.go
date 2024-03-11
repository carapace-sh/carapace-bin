package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registryCmd = &cobra.Command{
	Use:     "registry",
	Short:   "manage the flake registry",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registryCmd).Standalone()

	rootCmd.AddCommand(registryCmd)
}
