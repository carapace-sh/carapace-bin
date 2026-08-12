package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify repository state consistency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_verifyCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_verifyCmd)
}
