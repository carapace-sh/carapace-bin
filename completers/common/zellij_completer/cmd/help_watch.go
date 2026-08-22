package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a session (read-only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_watchCmd).Standalone()

	helpCmd.AddCommand(help_watchCmd)
}
