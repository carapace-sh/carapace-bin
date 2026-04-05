package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the new configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_buildCmd).Standalone()

	addOsRebuildFlags(os_buildCmd)

	osCmd.AddCommand(os_buildCmd)
}
