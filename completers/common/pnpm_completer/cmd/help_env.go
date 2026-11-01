package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage Node.js versions. Deprecated in favour of `pnpm runtime`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_envCmd).Standalone()

	helpCmd.AddCommand(help_envCmd)
}
