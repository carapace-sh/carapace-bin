package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commits on active branches in your workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_logCmd).Standalone()

	helpCmd.AddCommand(help_logCmd)
}
