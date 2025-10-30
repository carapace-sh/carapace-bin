package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "Commands for interacting with forges like GitHub, GitLab (coming soon), etc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_forgeCmd).Standalone()

	helpCmd.AddCommand(help_forgeCmd)
}
