package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_upgradeCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove stored offline transaction and delete cached package files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_upgradeCleanCmd).Standalone()

	system_upgradeCmd.AddCommand(system_upgradeCleanCmd)
}
