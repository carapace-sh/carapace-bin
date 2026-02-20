package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_expose_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add exposed binaries from an environment to your global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_expose_addCmd).Standalone()

	global_help_exposeCmd.AddCommand(global_help_expose_addCmd)
}
