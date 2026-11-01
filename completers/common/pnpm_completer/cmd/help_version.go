package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Apply the pending change intents (`pnpm version -r`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_versionCmd).Standalone()

	helpCmd.AddCommand(help_versionCmd)
}
