package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes dependencies from an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_removeCmd).Standalone()

	help_globalCmd.AddCommand(help_global_removeCmd)
}
