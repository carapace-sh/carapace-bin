package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_verify_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Verify repository state consistency (default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_verify_stateCmd).Standalone()

	repository_verify_stateCmd.Flags().Bool("heal", false, "Attempt to heal discrepancies found in a new staged state")
	repository_verify_stateCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_verify_stateCmd.Flags().String("path", "", "Optional path in the repository to start verification from")
	repository_verifyCmd.AddCommand(repository_verify_stateCmd)
}
