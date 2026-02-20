package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_expose_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add exposed binaries from an environment to your global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_expose_addCmd).Standalone()

	help_global_exposeCmd.AddCommand(help_global_expose_addCmd)
}
