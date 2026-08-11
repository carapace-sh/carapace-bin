package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify repository state consistency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_verifyCmd).Standalone()

	repository_verifyCmd.Flags().Bool("heal", false, "Attempt to heal discrepancies found in a new staged state")
	repository_verifyCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_verifyCmd.Flags().String("path", "", "Optional path in the repository to start verification from (for state verification)")
	repositoryCmd.AddCommand(repository_verifyCmd)
}
