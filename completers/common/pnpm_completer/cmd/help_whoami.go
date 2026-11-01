package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Displays your pnpm username",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_whoamiCmd).Standalone()

	helpCmd.AddCommand(help_whoamiCmd)
}
