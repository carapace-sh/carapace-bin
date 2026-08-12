package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify repository state consistency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_verifyCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_verifyCmd)
}
