package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup zellij and check its configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_setupCmd).Standalone()

	helpCmd.AddCommand(help_setupCmd)
}
