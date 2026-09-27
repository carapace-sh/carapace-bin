package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_setScriptCmd = &cobra.Command{
	Use:   "set-script",
	Short: "Set a script in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_setScriptCmd).Standalone()

	helpCmd.AddCommand(help_setScriptCmd)
}
