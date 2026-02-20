package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_expose_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add exposed binaries from an environment to your global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_expose_help_addCmd).Standalone()

	global_expose_helpCmd.AddCommand(global_expose_help_addCmd)
}
