package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Manage your scripts with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_scriptsCmd).Standalone()

	helpCmd.AddCommand(help_scriptsCmd)
}
