package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Check for outdated package and GitHub Actions dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_outdatedCmd).Standalone()

	helpCmd.AddCommand(help_outdatedCmd)
}
