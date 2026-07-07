package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to another branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_switchCmd).Standalone()

	helpCmd.AddCommand(help_switchCmd)
}
