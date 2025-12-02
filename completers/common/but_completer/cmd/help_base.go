package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_baseCmd = &cobra.Command{
	Use:   "base",
	Short: "Commands for managing the base target branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_baseCmd).Standalone()

	helpCmd.AddCommand(help_baseCmd)
}
