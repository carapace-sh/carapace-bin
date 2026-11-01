package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the config key to the value provided. Shorthand for `pnpm config set`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_setCmd).Standalone()

	helpCmd.AddCommand(help_setCmd)
}
