package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds dependencies to an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_addCmd).Standalone()

	help_globalCmd.AddCommand(help_global_addCmd)
}
