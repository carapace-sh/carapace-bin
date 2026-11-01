package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the config value for the provided key. Shorthand for `pnpm config get`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_getCmd).Standalone()

	helpCmd.AddCommand(help_getCmd)
}
